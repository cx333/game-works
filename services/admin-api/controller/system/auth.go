package system

import (
	"github.com/cx333/game-works/services/admin-api/model/sysModel"
	"github.com/cx333/game-works/services/admin-api/resource"
	"github.com/cx333/game-works/services/admin-api/server/system"
	"github.com/gin-gonic/gin"
)

/**
 * @Author: wgl
 * @Description: 登陆
 * @File: auth.go
 * @Version: 1.0.0
 * @Date: 2025/4/30 21:43
 */

var sysAuth system.AuthSvr

// LoginHandler 用户登录
func LoginHandler(ctx *gin.Context) {
	var req sysModel.Auth
	if err := ctx.ShouldBindJSON(&req); err != nil {
		resource.ErrBind(ctx)
		return
	}

	token, err := sysAuth.LoginImpl(&req)
	if err != nil || token == "" {
		resource.ErrorMsg(ctx, "登录失败，用户名或密码错误")
		return
	}
	resource.Success(ctx, gin.H{"token": token})
}

// LogoutHandler 退出登录
func LogoutHandler(ctx *gin.Context) {
	// 如需处理服务端 Session，可在此清理
	// 示例中仅返回成功
	resource.SuccessMsg(ctx, "登出成功")
}

// RegisterHandler 用户注册
func RegisterHandler(ctx *gin.Context) {
	var req sysModel.AuthRegister
	if err := ctx.ShouldBindJSON(&req); err != nil {
		resource.ErrBind(ctx)
		return
	}

	ok := sysAuth.RegisterImpl(&req)
	if !ok {
		resource.ErrorMsg(ctx, "注册失败，用户名可能已存在")
		return
	}
	resource.SuccessMsg(ctx, "注册成功")
}
