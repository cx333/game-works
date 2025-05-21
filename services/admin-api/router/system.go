package router

import (
	"github.com/cx333/game-works/services/admin-api/controller/Permissions"
	"github.com/cx333/game-works/services/admin-api/controller/organize"
	"github.com/cx333/game-works/services/admin-api/controller/system"
	"github.com/gin-gonic/gin"
)

/**
 * @Author: wgl
 * @Description: system 模块路由
 * @File: system.go
 * @Version: 1.0.0
 * @Date: 2025/4/30 21:40
 */

func RegisterSystemRoutes(group *gin.RouterGroup) {
	auth := group.Group("/system")
	{
		auth.POST("/logout", Permissions.LogoutHandler)
		auth.POST("/register", Permissions.RegisterHandler)
		auth.GET("/codes", Permissions.GetAccessCodes)
		auth.POST("/refresh", Permissions.RefreshToken)
	}
	user := group.Group("/user")
	{
		user.GET("/info", organize.GetUserInfo)
	}
	menu := group.Group("/menu")
	{
		menu.GET("/tree", system.GetAllMenus)
	}
}
