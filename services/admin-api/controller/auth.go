package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 登录
func Login(c *gin.Context) {
	// 这里只做演示，实际应校验用户名密码
	c.JSON(http.StatusOK, gin.H{
		"id":          1,
		"realName":    "Admin",
		"roles":       []string{"admin"},
		"username":    "admin",
		"homePath":    "/workspace",
		"accessToken": "mock-token",
	})
}

// 退出登录
func Logout(c *gin.Context) {
	c.Status(http.StatusOK)
}

// 获取权限码
func GetAccessCodes(c *gin.Context) {
	c.JSON(http.StatusOK, []string{"AC_100010", "AC_100020"})
}
