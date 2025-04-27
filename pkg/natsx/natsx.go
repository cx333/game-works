package natsx

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
)

var (
	ErrNotConnected = errors.New("nats: not connected to NATS server")
	ErrNilHandler   = errors.New("nats: message handler cannot be nil")
	ErrSubscription = errors.New("nats: subscription failed")
)

// NatsConn 封装 NATS 连接
type NatsConn struct {
	conn       *nats.Conn
	url        string
	opts       []nats.Option
	logger     *zap.Logger
	subs       map[string]*nats.Subscription
	subsMutex  sync.RWMutex
	quitChan   chan struct{}
	reconnects int
}

// NatsConfig 配置参数
type NatsConfig struct {
	URL            string
	Name           string
	MaxReconnects  int
	ReconnectWait  time.Duration
	ConnectTimeout time.Duration
	PingInterval   time.Duration
	MaxPingsOut    int
	Logger         *zap.Logger
}

// New 创建 NATS 连接
func New(config NatsConfig, opts ...nats.Option) (*NatsConn, error) {
	if config.Logger == nil {
		config.Logger = zap.NewNop()
	}

	// 默认配置
	defaultOpts := []nats.Option{
		nats.Name(config.Name),
		nats.MaxReconnects(config.MaxReconnects),
		nats.ReconnectWait(config.ReconnectWait),
		nats.Timeout(config.ConnectTimeout),
		nats.PingInterval(config.PingInterval),
		nats.MaxPingsOutstanding(config.MaxPingsOut),
		nats.DisconnectErrHandler(func(c *nats.Conn, err error) {
			config.Logger.Warn("NATS连接断开", zap.Error(err))
		}),
		nats.ReconnectHandler(func(c *nats.Conn) {
			config.Logger.Info("NATS重新连接成功", zap.String("url", c.ConnectedUrl()),
				zap.String("server", c.ConnectedServerId()))
		}),
		nats.ClosedHandler(func(c *nats.Conn) {
			config.Logger.Error("NATS连接永久关闭")
		}),
		nats.ErrorHandler(func(c *nats.Conn, sub *nats.Subscription, err error) {
			config.Logger.Error("NATS 异常", zap.Error(err))
		}),
		nats.DiscoveredServersHandler(func(c *nats.Conn) {
			config.Logger.Info("发现新服务器", zap.Strings("servers", c.DiscoveredServers()))
		}),
	}

	// 合并用户自定义配置
	finalOpts := append(defaultOpts, opts...)
	conn, err := nats.Connect(config.URL, finalOpts...)
	if err != nil {
		config.Logger.Error("Failed to connect to NATS", zap.Error(err))
		return nil, fmt.Errorf("natsx: failed to connect: %w", err)
	}
	nc := &NatsConn{
		conn:     conn,
		url:      config.URL,
		opts:     finalOpts,
		logger:   config.Logger,
		subs:     make(map[string]*nats.Subscription),
		quitChan: make(chan struct{}),
	}
	return nc, nil
}

// SubscribeWithRetry 带自动重试的订阅
func (c *NatsConn) SubscribeWithRetry(subject string, handler nats.MsgHandler) error {
	if handler == nil {
		return ErrNilHandler
	}

	go c.manageSubscription(subject, handler)
	return nil
}

func (c *NatsConn) manageSubscription(subject string, handler nats.MsgHandler) {
	retryDelay := time.Second
	maxRetryDelay := 30 * time.Second

	for {
		select {
		case <-c.quitChan:
			return
		default:
			if !c.IsConnected() {
				c.logger.Warn("NATS连接断开，等待重连...",
					zap.String("subject", subject),
					zap.Duration("retry_delay", retryDelay))
				time.Sleep(retryDelay)
				retryDelay = c.increaseDelay(retryDelay, maxRetryDelay)
				continue
			}
			sub, err := c.conn.Subscribe(subject, handler)
			if err != nil {
				c.logger.Error("订阅失败",
					zap.String("subject", subject),
					zap.Error(err))
				time.Sleep(retryDelay)
				retryDelay = c.increaseDelay(retryDelay, maxRetryDelay)
				continue
			}

			c.subsMutex.Lock()
			c.subs[subject] = sub
			c.subsMutex.Unlock()

			c.logger.Info("订阅成功", zap.String("subject", subject))
			retryDelay = time.Second // 重置重试间隔

			// 监控订阅状态
			ticker := time.NewTicker(5 * time.Second)
		monitorLoop:
			for {
				select {
				case <-c.quitChan:
					ticker.Stop()
					sub.Unsubscribe()
					return
				case <-ticker.C:
					if !sub.IsValid() {
						c.logger.Warn("订阅失效，重新订阅...",
							zap.String("subject", subject))
						ticker.Stop()
						break monitorLoop
					}
				}
			}
		}
	}
}

// Publish 发布消息
func (c *NatsConn) Publish(subject string, data []byte) error {
	if !c.IsConnected() {
		return ErrNotConnected
	}
	return c.conn.Publish(subject, data)
}

// Request 发送请求并等待响应
func (c *NatsConn) Request(subject string, data []byte, timeout time.Duration) (*nats.Msg, error) {
	if !c.IsConnected() {
		return nil, ErrNotConnected
	}
	return c.conn.Request(subject, data, timeout)
}

// Close 关闭连接和所有订阅
func (c *NatsConn) Close() {
	if !c.IsConnected() {
		return
	}
	close(c.quitChan)
	c.subsMutex.Lock()
	defer c.subsMutex.Unlock()

	for subject, sub := range c.subs {
		if err := sub.Unsubscribe(); err != nil {
			c.logger.Warn("取消订阅失败",
				zap.String("subject", subject),
				zap.Error(err))
		}
	}

	if c.conn != nil && !c.conn.IsClosed() {
		c.conn.Close()
	}
}

// IsConnected 检查连接状态
func (c *NatsConn) IsConnected() bool {
	return c.conn != nil && c.conn.IsConnected()
}

// Drain 优雅关闭连接
func (c *NatsConn) Drain() error {
	if c.conn == nil {
		return ErrNotConnected
	}
	return c.conn.Drain()
}

func (c *NatsConn) increaseDelay(current, max time.Duration) time.Duration {
	next := current * 2
	if next > max {
		return max
	}
	return next
}
