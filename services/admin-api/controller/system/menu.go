package system

import (
	"github.com/cx333/game-works/services/admin-api/resource"
	"github.com/cx333/game-works/services/admin-api/server/system"
	"github.com/gin-gonic/gin"
	"net/http"
)

var menuSvr = system.MenuSvr{}

func GetAllMenus(ctx *gin.Context) {
	userID, exists := ctx.Get("userID")
	if !exists {
		resource.ErrorCode(ctx, http.StatusUnauthorized, "User not authenticated")
		return
	}
	tree, err := menuSvr.MenuTree(userID.(uint))
	if err != nil {
		return
	}
	resource.Success(ctx, tree)
}
