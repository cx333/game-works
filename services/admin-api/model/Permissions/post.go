package Permissions

import (
	"github.com/cx333/game-works/services/admin-api/resource"
)

type PermPost struct {
	Name   string `gorm:"size:64;not null" json:"name"`
	Path   string `gorm:"size:64;not null" json:"path"`
	Method int    `json:"method"`
	Status int    `gorm:"default:1" json:"status"` // 0禁用, 1正常
	Desc   string `json:"desc"`
	resource.Models
}

func (p *PermPost) TableName() string {
	return "perm_post"
}
