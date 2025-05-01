package router

import (
	"github.com/cx333/game-works/services/admin-api/controller"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(r *gin.RouterGroup) {
	r.POST("/sys/login", controller.Login)
	r.POST("/auth/refresh", controller.RefreshToken)
	r.POST("/auth/logout", controller.Logout)
	r.GET("/auth/codes", controller.GetAccessCodes)
}
