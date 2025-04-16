package main

import (
	"github.com/cx333/game-works/pkg/logger"
	"github.com/cx333/game-works/services/gateway/transport"
)

/**
 * @Author: wgl
 * @Description:
 * @File: main
 * @Version: 1.0.0
 * @Date: 2025/4/16 20:31
 */

func main() {
	logger.Init("gateway", "debug", "./logs")
	transport.StartWebSocketServer(":9001")
}
