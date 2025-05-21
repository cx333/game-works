package main

import (
	"github.com/cx333/game-works/pkg/logger"
	"github.com/cx333/game-works/pkg/natsx"
	"github.com/cx333/game-works/pkg/shared"
	"github.com/cx333/game-works/services/gateway/transport"
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
		URL:            "nats://192.168.1.63:4222",
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
	shared.GatewayNats = nc
}

func main() {
	go transport.StartTcpServer()
	logger.Info("程序启动")
	// 阻塞主线程
	select {}
}
