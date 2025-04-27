package main

import (
	"fmt"
	"github.com/cx333/game-works/pkg/logger"
	"github.com/cx333/game-works/pkg/natsx"
	"github.com/cx333/game-works/services/gateway/shared"
	"github.com/cx333/game-works/services/gateway/transport"
	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
	"time"
)

/**
 * @Author: wgl
 * @Description:
 * @File: main
 * @Version: 1.0.0
 * @Date: 2025/4/16 20:31
 */

func init() {
	logger.Init("gateway", "debug", "./logs/")
	defer logger.Sync()
	config := natsx.NatsConfig{
		URL:            "nats://192.168.1.22:4222",
		Name:           "game-server",
		MaxReconnects:  -1, // 无限重连
		ReconnectWait:  2 * time.Second,
		ConnectTimeout: 5 * time.Second,
	}
	nc, err := natsx.New(config)
	if err != nil {
		logger.Warn("Failed to connect to NATS", err)
		return
	}
	shared.NatsConn = nc
}

func main() {
	go transport.StartTcpServer()
	logger.Info("程序启动")

	time.Sleep(time.Second * 2)

	err := shared.NatsConn.SubscribeWithRetry("match.request", func(m *nats.Msg) {
		fmt.Println("收到消息内容:", string(m.Data))
	})
	if err != nil {
		logger.Error("订阅失败", zap.Error(err))
		return
	}

	for i := 0; i < 10; i++ {
		err := shared.NatsConn.Publish("match.request", []byte(fmt.Sprintf("test %d", i)))
		if err != nil {
			logger.Error("发布失败", zap.Error(err))
		}
		time.Sleep(200 * time.Millisecond) // 稍微慢点
	}

	// 阻塞主线程
	select {}
}
