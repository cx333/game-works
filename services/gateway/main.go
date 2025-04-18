package main

import (
	"github.com/cx333/game-works/pkg/logger"
	"github.com/cx333/game-works/pkg/natsx"
	"github.com/cx333/game-works/services/gateway/transport"
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
	conn, err := natsx.New("nats://192.168.1.22:4222")
	if err != nil {
		logger.Error(err.Error())
	}
	NatsConn = *conn
}

func main() {
	transport.StartWebSocketServer(":9001")
}
