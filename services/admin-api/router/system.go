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
	auth := group.Group("/auth")
	{
		auth.POST("/logout", system.LogoutHandler)
		auth.POST("/register", system.RegisterHandler)
		auth.GET("/codes", system.GetAccessCodes)
		auth.POST("/refresh", system.RefreshToken)
	}
	user := group.Group("/user")
	{
		user.GET("/info", system.GetUserInfo)
	}
	menu := group.Group("/menu")
	{
		menu.GET("/tree", system.GetAllMenus)
	}
}
