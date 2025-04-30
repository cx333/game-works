package router

import (
	"github.com/cx333/game-works/services/admin-api/controller/system"
	"github.com/gin-gonic/gin"
)

/**
 * @Author: wgl
 * @Description: sysModel 模块路由
 * @File: sysModel.go
 * @Version: 1.0.0
 * @Date: 2025/4/30 21:40
 */

func RegisterSystemRoutes(group *gin.RouterGroup) {
	sys := group.Group("/sys")
	{
		sys.POST("/login", system.LoginHandler)
		sys.POST("/logout", system.LogoutHandler)
		sys.POST("/register", system.RegisterHandler)
	}
}
