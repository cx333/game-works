package Permissions

import (
	"github.com/cx333/game-works/services/admin-api/resource"
)

type PermCode struct {
	Name string `json:"name"`
	Code string `gorm:"size:64;not null" json:"code"`
	Desc string `gorm:"size:256" json:"desc"`
	resource.Models
}

func (*PermCode) TableName() string {
	return "perm_code"
}
