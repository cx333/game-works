package middleware

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/cx333/game-works/pkg/logger"
	"github.com/cx333/game-works/services/admin-api/model"
)

/**
 * @Author: wgl
 * @Description:
 * @File: casbin
 * @Version: 1.0.0
 * @Date: 2025/4/30 19:59
 */

func SetupCasbin() *casbin.Enforcer {
	a, _ := gormadapter.NewAdapterByDB(model.PgsqlDB)
	e, _ := casbin.NewEnforcer("config/casbin_model.conf", a)
	err := e.LoadPolicy()
	if err != nil {
		logger.Error("LoadPolicy err", err.Error())
		return nil
	}
	return e
}
