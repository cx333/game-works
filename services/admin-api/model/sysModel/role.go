package sysModel

import "time"

type SysRole struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"uniqueIndex;size:64;not null" json:"name"`
	Label     string    `gorm:"size:64;not null" json:"label"`
	Desc      string    `gorm:"size:256" json:"desc"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Users     []SysUser `gorm:"many2many:user_roles;" json:"users"`
	Codes     []SysCode `gorm:"many2many:role_codes;" json:"codes"`
}

func (r *SysRole) TableName() string {
	return "sys_role"
}
