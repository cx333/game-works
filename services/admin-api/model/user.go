package model

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Username  string    `gorm:"uniqueIndex;size:64;not null" json:"username"`
	Password  string    `gorm:"size:128;not null" json:"-"`
	RealName  string    `gorm:"size:64;not null" json:"realName"`
	Avatar    string    `gorm:"size:256" json:"avatar"`
	HomePath  string    `gorm:"size:128" json:"homePath"`
	Desc      string    `gorm:"size:256" json:"desc"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Roles     []Role    `gorm:"many2many:user_roles;" json:"roles"`
}
