package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

/**
 * @Author: wgl
 * @Description:
 * @File: jwt
 * @Version: 1.0.0
 * @Date: 2025/4/30 21:57
 */

// JWTAuthMiddleware 校验 jwt
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "请登录"})
			c.Abort()
			return
		}
		token := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := ParseToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "无效令牌"})
			c.Abort()
			return
		}
		c.Set("userID", claims.UserID)
		c.Set("username", claims.Username)
		c.Next()
	}
}
