package system

import (
	"github.com/cx333/game-works/services/admin-api/middleware"
	"github.com/cx333/game-works/services/admin-api/model"
	"github.com/cx333/game-works/services/admin-api/model/sysModel"
	"golang.org/x/crypto/bcrypt"
)

/**
 * @Author: wgl
 * @Description:
 * @File: auth_impl
 * @Version: 1.0.0
 * @Date: 2025/4/30 22:11
 */

type AuthImpl interface {
	LoginImpl(req *sysModel.Auth) (token string, err error)
	LogoutImpl(req *sysModel.Auth) bool
	RegisterImpl(req *sysModel.AuthRegister) bool
}

type AuthSvr struct{}

func (a AuthSvr) LoginImpl(req *sysModel.Auth) (token string, err error) {
	// 实现登录逻辑
	generateToken, err := middleware.GenerateToken(req.ID, req.Username)
	if err != nil {
		return "", err
	}
	return generateToken, err
}

func (a AuthSvr) LogoutImpl(req *sysModel.Auth) bool {
	return true
}

func (a AuthSvr) RegisterImpl(req *sysModel.AuthRegister) bool {
	var count int64
	err := model.PgsqlDB.Model(&sysModel.SysUser{}).Where("username = ?", req.Username).Count(&count).Error
	if err != nil || count > 0 {
		return false
	}
	// 密码加密
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return false
	}

	newUser := sysModel.SysUser{
		Username: req.Username,
		Password: string(hashedPwd),
		Nickname: req.Nickname,
		Email:    req.Email,
		Phone:    req.Phone,
		Status:   1,
		DeptID:   req.DeptId,
		PostID:   req.PostId,
		Avatar:   req.Avatar,
		Remark:   req.Remark,
	}
	if err := model.PgsqlDB.Create(&newUser).Error; err != nil {
		return false
	}
	return true
}

// 强制保证 AuthSvr 实现 AuthImpl 接口
var _ AuthImpl = AuthSvr{}
