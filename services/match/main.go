package main

import (
	"github.com/cx333/game-works/pkg/logger"
	"github.com/cx333/game-works/pkg/natsx"
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
	conn, err := natsx.New("nats://192.168.1.22:4222")
	if err != nil {
		logger.Error(err.Error())
	}
	NatsConn = *conn

}

func main() {

	natsx.RegisterTopic(natsx.MatchRequestTopic, "玩家发起匹配请求")
	natsx.RegisterTopic(natsx.MatchResultTopic, "服务返回匹配成功的结果")

	natsx.PrintRegisteredTopics("match")

}
