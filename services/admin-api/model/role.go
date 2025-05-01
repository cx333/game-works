package model

import (
	"github.com/cx333/game-works/services/admin-api/model/sysModel"
	"time"
)

type Role struct {
	ID        uint            `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string          `gorm:"uniqueIndex;size:64;not null" json:"name"`
	Label     string          `gorm:"size:64;not null" json:"label"`
	Desc      string          `gorm:"size:256" json:"desc"`
	CreatedAt time.Time       `json:"createdAt"`
	UpdatedAt time.Time       `json:"updatedAt"`
	Users     []User          `gorm:"many2many:user_roles;" json:"users"`
	Codes     []sysModel.Code `gorm:"many2many:role_codes;" json:"codes"`
}
