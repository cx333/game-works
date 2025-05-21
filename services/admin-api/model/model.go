package model

import (
	"github.com/cx333/game-works/pkg/db"
	"github.com/cx333/game-works/pkg/logger"
	"github.com/cx333/game-works/services/admin-api/config"
	"github.com/cx333/game-works/services/admin-api/model/Permissions"
	"github.com/cx333/game-works/services/admin-api/model/organize"
	"github.com/cx333/game-works/services/admin-api/model/system"
	"gorm.io/gorm"
)

var PgsqlDB *gorm.DB

//var tenantDBs = sync.Map{}

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
		&organize.User{},
		&organize.Dept{},
		&organize.Company{},

		&Permissions.PermRole{},
		&Permissions.PermCasbinRule{},
		&Permissions.PermPost{},
		&Permissions.PermCode{},
		&Permissions.PermUserRole{},
		&Permissions.PermRoleDataScope{},

		&system.SysMenu{},
		&system.SysDictType{},
		&system.SysDictData{},
	)
	if err != nil {
		logger.Error("❌ 初始化数据表失败", err)
		return
	}
	logger.Debug("✅ 初始化数据表成功")
}

//func CreateTenant(tenantName string) error {
//	dbName := "tenant_" + strings.ToLower(tenantName)
//
//	// Step 1: 在主库插入租户记录
//	err := PgsqlDB.Create(&SysTenant{Name: tenantName, DbName: dbName}).Error
//	if err != nil {
//		return err
//	}
//
//	// Step 2: 创建新数据库
//	err = db.Exec("CREATE DATABASE " + dbName).Error
//	if err != nil {
//		return err
//	}
//
//	// Step 3: 初始化新数据库连接
//	newConf := db.PgsqlConfig{... DbName: dbName }
//	newDB, err := db.InitPostgres(&newConf)
//	if err != nil {
//		return err
//	}
//
//	// Step 4: 自动迁移表结构
//	err = newDB.AutoMigrate(
//		&organize.SysUser{},
//		&Permissions.SysRole{},
//		&system.SysMenu{},
//		&organize.SysDept{},
//		&Permissions.SysPost{},
//		&system.SysDictType{},
//		&system.SysDictData{},
//	)
//	if err != nil {
//		return err
//	}
//
//	// Step 5: 缓存连接
//	tenantDBs.Store(dbName, newDB)
//	return nil
//}

//func GetTenantDB(tenantDbName string) *gorm.DB {
//	db, ok := tenantDBs.Load(tenantDbName)
//	if ok {
//		return db.(*gorm.DB)
//	}
//	// fallback: 从配置重新加载连接
//	return nil
//}
