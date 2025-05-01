package sysModel

import "time"

type SysMenu struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"size:64;not null" json:"name"`
	Path      string    `gorm:"size:128;not null" json:"path"`
	Component string    `gorm:"size:128" json:"component"`
	ParentID  *uint     `gorm:"index" json:"parentId"`
	Order     int       `gorm:"default:0" json:"order"`
	Children  []SysMenu `gorm:"foreignKey:ParentID" json:"children"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (m *SysMenu) TableName() string {
	return "sys_menu"
}
