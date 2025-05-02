package system

import (
	"github.com/cx333/game-works/services/admin-api/model"
	"github.com/cx333/game-works/services/admin-api/model/sysModel"
)

type UserImpl interface {
	GetUserInfo(userId uint) (*sysModel.UserInfoRes, error)
}

type UserSvr struct {
}

func (u UserSvr) GetUserInfo(userId uint) (*sysModel.UserInfoRes, error) {
	res := &sysModel.UserInfoRes{}
	err := model.PgsqlDB.Model(&sysModel.SysUser{}).
		Where("id = ?", userId).
		First(res).Error

	return res, err
}

var _ UserImpl = UserSvr{}
