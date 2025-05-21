package system

import (
	"github.com/cx333/game-works/services/admin-api/model"
	"github.com/cx333/game-works/services/admin-api/model/system"
)

type MenuImpl interface {
	MenuTree(uint) ([]system.RouteRecordStringComponent, error)
}

type MenuSvr struct{}

// MenuTree 获取菜单树（支持多个顶级菜单）
func (m MenuSvr) MenuTree(userId uint) ([]system.RouteRecordStringComponent, error) {
	menus, err := m.GetAllMenus()
	if err != nil {
		return nil, err
	}
	return buildRouteTree(menus), nil
}

func (m MenuSvr) GetAllMenus() ([]system.SysMenu, error) {
	var menus []system.SysMenu
	err := model.PgsqlDB.
		Model(&system.SysMenu{}).
		Where("deleted_at IS NULL").
		Order("\"order\" ASC").
		Find(&menus).Error
	return menus, err
}

func buildRouteTree(menus []system.SysMenu) []system.RouteRecordStringComponent {
	nodeMap := make(map[uint64]*system.RouteRecordStringComponent)
	var roots []*system.RouteRecordStringComponent // 改为指针列表

	// 第一步：创建所有节点（初始化 Children）
	for _, menu := range menus {
		nodeMap[menu.ID] = &system.RouteRecordStringComponent{
			Name:      menu.Name,
			Path:      menu.Path,
			Component: menu.Component,
			Redirect:  menu.Redirect,
			Meta: &system.RouteMeta{
				Title:           menu.Title,
				Icon:            menu.Icon,
				Order:           menu.Order,
				AffixTab:        menu.AffixTab,
				HideInMenu:      menu.HideInMenu,
				KeepAlive:       menu.KeepAlive,
				OpenInNewWindow: menu.OpenInNewWin,
				NoBasicLayout:   menu.NoBasicLayout,
			},
			Children: make([]system.RouteRecordStringComponent, 0), // 显式初始化
		}
	}

	// 第二步：构建树结构
	for _, menu := range menus {
		node := nodeMap[menu.ID]
		if menu.ParentID == 0 {
			roots = append(roots, node) // 直接存储指针
		} else {
			if parent, ok := nodeMap[menu.ParentID]; ok {
				parent.Children = append(parent.Children, *node) // 解引用添加
			}
		}
	}

	// 转换为值类型返回（避免外部修改内部数据）
	result := make([]system.RouteRecordStringComponent, len(roots))
	for i, root := range roots {
		result[i] = *root
	}
	return result
}

// 强制保证 AuthSvr 实现 AuthImpl 接口
var _ MenuImpl = MenuSvr{}
