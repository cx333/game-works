package main

import (
	"github.com/cx333/game-works/pkg/logger"
	"github.com/cx333/game-works/services/admin-api/config"
	"github.com/cx333/game-works/services/admin-api/middleware"
	"github.com/cx333/game-works/services/admin-api/model"
	"github.com/cx333/game-works/services/admin-api/router"
	"github.com/gin-gonic/gin"
)

/**
 * @Author: wgl
 * @Description: 管理系统
 * @File: main
 * @Version: 1.0.0
 * @Date: 2025/4/30 18:52
 */

func init() {
	config.InitConfig()
	logger.Init("admin", logger.DebugLevel, "./logs")
	defer logger.Sync()
	model.InitModel()
	middleware.SetupCasbin()
}

func main() {
	r := gin.Default()
	router.RegisterRoutes(r)
	err := r.Run(":8080")
	if err != nil {
		logger.Error("admin 启动失败", err.Error())
		return
	}
}
