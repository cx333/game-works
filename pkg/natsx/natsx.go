package natsx

import (
	"errors"
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
)

var (
	ErrNotConnected = errors.New("nats: not connected to NATS server")
	ErrNilHandler   = errors.New("nats: message handler cannot be nil")
)

// NatsConn 封装 NATS 连接
type NatsConn struct {
	conn *nats.Conn
}

// NatsConnInterface 定义 NATS 客户端接口
type NatsConnInterface interface {
	Publish(subject string, data []byte) error
	Subscribe(subject string, handler nats.MsgHandler) (*nats.Subscription, error)
	Close()
	IsConnected() bool
}

// New 创建 NATS 连接
func New(natsURL string, opts ...nats.Option) (*NatsConn, error) {
	// 默认配置
	defaultOpts := []nats.Option{
		nats.MaxReconnects(10),
		nats.ReconnectWait(2 * time.Second),
		nats.Name("game-server"),
		nats.Timeout(5 * time.Second),
		nats.PingInterval(20 * time.Second),
		nats.MaxPingsOutstanding(3),
	}

	// 合并用户自定义配置
	finalOpts := append(defaultOpts, opts...)

	conn, err := nats.Connect(natsURL, finalOpts...)
	if err != nil {
		return nil, fmt.Errorf("natsx: failed to connect: %w", err)
	}
	return &NatsConn{conn: conn}, nil
}

// Publish 发布消息
func (c *NatsConn) Publish(subject string, data []byte) error {
	if !c.IsConnected() {
		return ErrNotConnected
	}
	return c.conn.Publish(subject, data)
}

// Subscribe 订阅主题
func (c *NatsConn) Subscribe(subject string, handler nats.MsgHandler) (*nats.Subscription, error) {
	if !c.IsConnected() {
		return nil, ErrNotConnected
	}
	return c.conn.Subscribe(subject, handler)
}

// Close 关闭连接
func (c *NatsConn) Close() {
	if c.conn != nil && !c.conn.IsClosed() {
		c.conn.Close()
	}
}

// IsConnected 检查连接状态
func (c *NatsConn) IsConnected() bool {
	return c.conn != nil && c.conn.IsConnected()
}

// RouterRegister 注册服务
func (c *NatsConn) RouterRegister(command string, fc func()) {

}
