package system

import (
	"github.com/cx333/game-works/services/admin-api/resource"
)

type SysDictType struct {
	Name   string `gorm:"size:64;not null" json:"name"`
	Type   string `gorm:"size:64;not null;uniqueIndex" json:"type"`
	Status int    `gorm:"default:1" json:"status"`
	Remark string `json:"remark"`
	resource.Models
}

type SysDictData struct {
	TypeID uint64 `json:"type_id"`
	Label  string `gorm:"size:64;not null" json:"label"`
	Value  string `gorm:"size:128;not null" json:"value"`
	Sort   int    `gorm:"default:0" json:"sort"`
	Status int    `gorm:"default:1" json:"status"`
	Remark string `json:"remark"`
	resource.Models
}

func (d *SysDictType) TableName() string {
	return "sys_dict_type"
}

func (d *SysDictData) TableName() string {
	return "sys_dict_data"
}
