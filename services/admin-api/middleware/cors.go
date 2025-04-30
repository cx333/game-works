package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
 * @Author: wgl
 * @Description: 跨域
 * @File: cors
 * @Version: 1.0.0
 * @Date: 2025/4/30 21:50
 */

// CORSMiddleware 允许所有跨域请求（你也可以按需配置）
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 设置跨域响应头
		c.Header("Access-Control-Allow-Origin", "*") // ⚠️ 如果是生产环境，建议替换为指定域名
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With, Token")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")

		// 处理预检请求
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
