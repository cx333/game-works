package system

import (
	"github.com/cx333/game-works/services/admin-api/model/sysModel"
)

type MenuImpl interface {
	MenuTree(uint) (*sysModel.RouteRecordStringComponent, error)
}

type MenuSvr struct{}

// MenuTree 获取菜单
func (m MenuSvr) MenuTree(userId uint) (*sysModel.RouteRecordStringComponent, error) {
	//var list = &sysModel.SysMenu{}
	//model.PgsqlDB.Model(&sysModel.SysMenu{}).Where()

	return nil, nil

}

// 强制保证 AuthSvr 实现 AuthImpl 接口
var _ MenuImpl = MenuSvr{}
