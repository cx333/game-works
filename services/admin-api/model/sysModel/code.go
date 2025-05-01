package sysModel

import (
	"github.com/cx333/game-works/services/admin-api/model"
	"time"
)

type SysCode struct {
	ID        uint         `gorm:"primaryKey;autoIncrement" json:"id"`
	Code      string       `gorm:"uniqueIndex;size:64;not null" json:"code"`
	Desc      string       `gorm:"size:256" json:"desc"`
	CreatedAt time.Time    `json:"createdAt"`
	UpdatedAt time.Time    `json:"updatedAt"`
	Roles     []model.Role `gorm:"many2many:role_codes;" json:"roles"`
}

func (SysCode) TableName() string {
	return "sys_code"
}
