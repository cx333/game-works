package sysModel

import "time"

type SysMenu struct {
	ID        uint64    `gorm:"primaryKey" json:"id"`
	ParentID  uint64    `gorm:"default:0" json:"parent_id"`
	Title     string    `gorm:"size:128;not null" json:"title"`
	Name      string    `gorm:"size:64" json:"name"`       // 前端组件名
	Path      string    `gorm:"size:128" json:"path"`      // 路由路径
	Component string    `gorm:"size:128" json:"component"` // Vue/Tauri 组件路径
	Icon      string    `gorm:"size:64" json:"icon"`
	Sort      int       `gorm:"default:0" json:"sort"`
	Type      string    `gorm:"size:32" json:"type"` // dir/menu/button
	Visible   bool      `gorm:"default:true" json:"visible"`
	Perms     string    `gorm:"size:128" json:"perms"` // 权限标识
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (m *SysMenu) TableName() string {
	return "sys_menu"
}
