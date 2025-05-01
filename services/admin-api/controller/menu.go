package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 获取所有菜单
func GetAllMenus(c *gin.Context) {
	c.JSON(http.StatusOK, []gin.H{
		{
			"component": "BasicLayout",
			"meta":      gin.H{"order": -1, "title": "page.dashboard.title"},
			"name":      "Dashboard",
			"path":      "/",
			"redirect":  "/analytics",
			"children": []gin.H{
				{
					"name":      "Analytics",
					"path":      "/analytics",
					"component": "/dashboard/analytics/index",
					"meta":      gin.H{"affixTab": true, "title": "page.dashboard.analytics"},
				},
			},
		},
	})
}
