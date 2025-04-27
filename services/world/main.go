package main

import (
	"github.com/cx333/game-works/pkg/logger"
	"github.com/cx333/game-works/pkg/natsx"
	"github.com/cx333/game-works/services/gateway/router"
	"time"
)

/**
 * @Author: wgl
 * @Description:
 * @File: main
 * @Version: 1.0.0
 * @Date: 2025/4/21 22:14
 */

var NatsConn = natsx.NatsConn{}

func init() {
	logger.Init("world", logger.DebugLevel, "../logs/world")
	defer logger.Sync()
	conn, err := natsx.New("nats://192.168.1.22:4222", nil)
	if err != nil {
		logger.Error(err.Error())
	}
	NatsConn = *conn

	router.Register("world", matchRequestHandler)
	retryDelay := time.Second // 初始重试间隔
	maxRetryDelay := 30 * time.Second
}

func main() {

}
