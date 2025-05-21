package organize

import (
	"github.com/cx333/game-works/services/admin-api/model"
	"github.com/cx333/game-works/services/admin-api/model/organize"
)

type UserImpl interface {
	GetUserInfo(userId uint) (*organize.UserInfoRes, error)
}

type UserSvr struct {
}

func (u UserSvr) GetUserInfo(userId uint) (*organize.UserInfoRes, error) {
	res := organize.UserInfoRes{}
	err := model.PgsqlDB.Model(&organize.User{}).
		Where("id = ?", userId).
		Select("avatar, real_name, roles, id as user_id, username").
		First(&res).Error
	return &res, err
}

var _ UserImpl = UserSvr{}
