package router

import (
	"github.com/cx333/game-works/services/admin-api/controller"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.RouterGroup) {
	r.GET("/info", controller.GetUserInfo)
}
