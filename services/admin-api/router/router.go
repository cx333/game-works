package router

import (
	"github.com/cx333/game-works/services/admin-api/middleware"
	"github.com/gin-gonic/gin"
)

/**
 * @Author: wgl
 * @Description: 主路由
 * @File: router.go
 * @Version: 1.0.0
 * @Date: 2025/4/30 21:40
 */

// RegisterRoutes 注册所有模块的路由
func RegisterRoutes(r *gin.Engine) {
	// 全局中间件（如 CORS、日志等）
	r.Use(middleware.CORSMiddleware())

	// API 分组（可按版本号扩展）
	api := r.Group("/api")

	// 系统基础模块（用户、权限、菜单等）
	RegisterSystemRoutes(api)

	// 插件模块路由注册（例如商城、报表等）
	//RegisterPluginRoutes(api)

	// CMS 模块路由
	//RegisterCMSRoutes(api)

	// OA 协同办公路由
	//RegisterOARoutes(api)

	// 其他模块...
}

func RegisterSystemRoutes(api *gin.RouterGroup) {
	RegisterAuthRoutes(api)
	RegisterUserRoutes(api.Group("/user"))
	RegisterMenuRoutes(api.Group("/menu"))
}
