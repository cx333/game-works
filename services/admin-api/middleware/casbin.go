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

func SetupCasbin() (*casbin.Enforcer, error) {
	a, err := gormadapter.NewAdapterByDB(model.PgsqlDB)
	if err != nil {
		logger.Error("Failed to create Casbin adapter:", err)
		return nil, err
	}

	e, err := casbin.NewEnforcer("config/casbin.conf", a)
	if err != nil {
		logger.Error("Failed to create Casbin enforcer:", err)
		return nil, err
	}

	if err := e.LoadPolicy(); err != nil {
		logger.Error("Failed to load Casbin policy:", err)
		return nil, err
	}

	return e, nil
}
