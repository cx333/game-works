package organize

import (
	"github.com/cx333/game-works/services/admin-api/resource"
	"github.com/cx333/game-works/services/admin-api/server/organize"
	"github.com/gin-gonic/gin"
	"net/http"
)

var sysUser = organize.UserSvr{}

// GetUserInfo 获取用户信息
func GetUserInfo(ctx *gin.Context) {
	userID, exists := ctx.Get("userID")
	if !exists {
		resource.ErrorCode(ctx, http.StatusUnauthorized, "User not authenticated")
		return
	}
	info, err := sysUser.GetUserInfo(userID.(uint))
	if err != nil {
		resource.Error(ctx)
		return
	}
	resource.Success(ctx, info)
}
