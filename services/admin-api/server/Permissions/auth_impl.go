package Permissions

import (
	"github.com/cx333/game-works/services/admin-api/middleware"
	"github.com/cx333/game-works/services/admin-api/model"
	"github.com/cx333/game-works/services/admin-api/model/organize"
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
	LoginImpl(req *organize.Auth) (token string, err error)
	LogoutImpl(req *organize.Auth) bool
	RegisterImpl(req *organize.AuthRegister) bool
}

type AuthSvr struct{}

func (a AuthSvr) LoginImpl(req *organize.Auth) (token string, err error) {
	var user organize.User
	err = model.PgsqlDB.Model(&organize.User{}).Where("username = ?", req.Username).First(&user).Error
	if err != nil || user.Username != req.Username {
		return "", err
	}
	// 2. 比对密码（核心逻辑）
	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password), // 数据库存储的哈希密码
		[]byte(req.Password),  // 用户输入的明文密码
	)
	if err != nil {
		return "", err
	}
	// 实现登录逻辑
	generateToken, err := middleware.GenerateToken(uint(user.ID), req.Username)
	if err != nil {
		return "", err
	}
	return generateToken, err
}

func (a AuthSvr) LogoutImpl(req *organize.Auth) bool {
	return true
}

func (a AuthSvr) RegisterImpl(req *organize.AuthRegister) bool {
	var count int64
	err := model.PgsqlDB.Model(&organize.User{}).Where("username = ?", req.Username).Count(&count).Error
	if err != nil || count > 0 {
		return false
	}
	// 密码加密
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return false
	}

	newUser := organize.User{
		Username: req.Username,
		Password: string(hashedPwd),
		Nickname: req.Nickname,
		Email:    req.Email,
		Phone:    req.Phone,
		Status:   0,
		Avatar:   req.Avatar,
		Desc:     req.Desc,
	}
	if err := model.PgsqlDB.Create(&newUser).Error; err != nil {
		return false
	}
	return true
}

// 强制保证 AuthSvr 实现 AuthImpl 接口
var _ AuthImpl = AuthSvr{}
