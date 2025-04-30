package main

import (
	"github.com/cx333/game-works/pkg/logger"
	"github.com/cx333/game-works/pkg/natsx"
	"github.com/cx333/game-works/services/chat/message"
	"github.com/cx333/game-works/services/chat/shared"
	"time"
)

/**
 * @Author: wgl
 * @Description: 聊天
 * @File: main
 * @Version: 1.0.0
 * @Date: 2025/4/22 21:40
 */

func init() {
	logger.Init("chat", "debug", "./logs/")
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
	shared.NatsConn = nc
}

func main() {
	message.SubMessage()
	select {}
}
