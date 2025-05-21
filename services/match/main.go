package main

import (
	"github.com/cx333/game-works/pkg/logger"
	"github.com/cx333/game-works/pkg/natsx"
	"github.com/cx333/game-works/pkg/shared"
	"time"
)

/**
 * @Author: wgl
 * @Description:
 * @File: main
 * @Version: 1.0.0
 * @Date: 2025/4/16 22:07
 */

var NatsConn = natsx.NatsConn{}

func init() {
	logger.Init("match", logger.DebugLevel, "./logs")
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
	shared.MatchNats = nc

}

func main() {

	natsx.RegisterTopic(natsx.MatchRequestTopic, "玩家发起匹配请求")
	natsx.RegisterTopic(natsx.MatchResultTopic, "服务返回匹配成功的结果")

	natsx.PrintRegisteredTopics("match")

}
