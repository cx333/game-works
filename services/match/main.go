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

func main() {
	logger.Init("match", "debug", "./logs")
	defer logger.Sync()
	natsx.RegisterTopic(natsx.MatchRequestTopic, "玩家发起匹配请求")
	natsx.RegisterTopic(natsx.MatchResultTopic, "服务返回匹配成功的结果")

	natsx.PrintRegisteredTopics("match")
	logger.Info("匹配服务启动成功")

}
