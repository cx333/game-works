package system

import (
	"github.com/cx333/game-works/services/admin-api/resource"
	"github.com/cx333/game-works/services/admin-api/server/system"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RouteMeta struct {
	Title           string   `json:"title"`
	Icon            string   `json:"icon,omitempty"`
	Order           int      `json:"order,omitempty"`
	AffixTab        bool     `json:"affixTab,omitempty"`
	Authority       []string `json:"authority,omitempty"`
	HideInMenu      bool     `json:"hideInMenu,omitempty"`
	KeepAlive       bool     `json:"keepAlive,omitempty"`
	OpenInNewWindow bool     `json:"openInNewWindow,omitempty"`
	NoBasicLayout   bool     `json:"noBasicLayout"` // 部分特殊页面如果不需要基础布局（页面顶部和侧边栏），可将noBasicLayout设置为true
}

type RouteRecordStringComponent struct {
	Name      string                       `json:"name"`
	Path      string                       `json:"path"`
	Component string                       `json:"component,omitempty"`
	Redirect  string                       `json:"redirect,omitempty"`
	Meta      *RouteMeta                   `json:"meta,omitempty"`
	Children  []RouteRecordStringComponent `json:"children,omitempty"`
}

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
