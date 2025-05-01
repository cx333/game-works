package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 获取用户信息
func GetUserInfo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"id":       1,
		"realName": "Admin",
		"roles":    []string{"admin"},
		"username": "admin",
		"homePath": "/workspace",
	})
}
