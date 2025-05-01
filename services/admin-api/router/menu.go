package router

import (
	"github.com/cx333/game-works/services/admin-api/controller"
	"github.com/gin-gonic/gin"
)

func RegisterMenuRoutes(r *gin.RouterGroup) {
	r.GET("/all", controller.GetAllMenus)
}
