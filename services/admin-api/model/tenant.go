package model

import "github.com/cx333/game-works/services/admin-api/resource"

// 多租户数据库

type SysTenant struct {
	Name   string
	DbName string // 租户数据库名
	Remark string
	resource.Models
}
