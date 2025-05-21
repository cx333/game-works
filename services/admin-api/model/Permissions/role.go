package Permissions

import "github.com/cx333/game-works/services/admin-api/resource"

type PermRole struct {
	Name string `gorm:"size:64;not null" json:"name"`
	Key  string `gorm:"size:64;not null" json:"key"`
	Desc string `gorm:"size:256" json:"desc"`
	resource.Models
}

func (r *PermRole) TableName() string {
	return "perm_role"
}
