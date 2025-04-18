package main

import (
	"github.com/cx333/game-works/pkg/logger"
	"github.com/cx333/game-works/pkg/natsx"
	"github.com/cx333/game-works/pkg/proto"
	"github.com/cx333/game-works/services/gateway/router"
	"github.com/cx333/game-works/services/gateway/transport"
	"github.com/nats-io/nats.go"
	protobuf "google.golang.org/protobuf/proto"
	"time"
)

/**
 * @Author: wgl
 * @Description:
 * @File: main
 * @Version: 1.0.0
 * @Date: 2025/4/16 20:31
 */

var NatsConn = natsx.NatsConn{}

func init() {
	logger.Init("gateway", "debug", "./logs")
	defer logger.Sync()
	conn, err := natsx.New("nats://192.168.1.22:4222", nil)
	if err != nil {
		logger.Error(err.Error())
	}
	NatsConn = *conn

	router.Register("gateway", matchRequestHandler)
	retryDelay := time.Second // 初始重试间隔
	maxRetryDelay := 30 * time.Second
	// 监听匹配请求
	go func() {
		for {
			// 检查连接是否有效
			if conn == nil || !conn.IsConnected() {
				logger.Error("NATS 连接已断开")
				time.Sleep(retryDelay)
				retryDelay = increaseDelay(retryDelay, maxRetryDelay)
				continue
			}
			sub, err := conn.Subscribe(natsx.MatchRequestTopic, onMatchRequest)
			if err != nil {
				logger.Error("匹配服务订阅失败:", err)
				time.Sleep(retryDelay)
				retryDelay = increaseDelay(retryDelay, maxRetryDelay)
				continue
			}

			logger.Info("匹配服务订阅成功，等待消息...")

			// 阻塞直到订阅出错（如连接断开）
			select {
			case <-time.After(retryDelay): // 定期检查订阅状态
				if !sub.IsValid() {
					logger.Warn("订阅已失效，重新订阅...")
					break
				}
			}
		}
	}()

}

// 指数退避计算下一次重试间隔
func increaseDelay(current, max time.Duration) time.Duration {
	next := current * 2
	if next > max {
		return max
	}
	return next
}

// 消息处理函数
func onMatchRequest(m *nats.Msg) {
	var req proto.MatchRequest
	if err := protobuf.Unmarshal(m.Data, &req); err != nil {
		logger.Error("反序列化 MatchRequest 失败:", err)
		return
	}

	logger.Debug("匹配服务处理请求:", req.GameMode, req.PlayerId)
}

func main() {
	transport.StartWebSocketServer(":9001")

}
