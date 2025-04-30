package model

import (
	"github.com/cx333/game-works/pkg/db"
	"github.com/cx333/game-works/pkg/logger"
	"github.com/cx333/game-works/services/admin-api/config"
	"github.com/cx333/game-works/services/admin-api/model/sysModel"
	"gorm.io/gorm"
)

var PgsqlDB *gorm.DB

func InitModel() {
	pgConf := db.PgsqlConfig{
		Host:     config.Config.Database.Pgsql.Host,
		Port:     config.Config.Database.Pgsql.Port,
		User:     config.Config.Database.Pgsql.User,
		Password: config.Config.Database.Pgsql.Password,
		DbName:   config.Config.Database.Pgsql.DbName,
		SslMode:  config.Config.Database.Pgsql.SslMode,
	}
	pgDb, err := db.InitPostgres(&pgConf)
	if err != nil {
		logger.Error("pgsql 连接失败", err)
		return
	}
	PgsqlDB = pgDb
	autoMigrate()

}

func autoMigrate() {
	err := PgsqlDB.AutoMigrate(
		&sysModel.SysUser{},
		&sysModel.SysRole{},
		&sysModel.SysMenu{},
		&sysModel.SysDept{},
		&sysModel.SysPost{},
		&sysModel.SysDictType{},
		&sysModel.SysDictData{},
		&sysModel.CasbinRule{},
	)
	if err != nil {
		logger.Error("❌ 初始化数据表失败", err)
		return
	}
	logger.Debug("✅ 初始化数据表成功")
}
