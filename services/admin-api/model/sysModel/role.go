package sysModel

import "time"

type SysRole struct {
	ID        uint64    `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"size:64;not null" json:"name"`
	Code      string    `gorm:"size:64;uniqueIndex;not null" json:"code"` // 如 admin, user
	Status    int       `gorm:"default:1" json:"status"`
	DataScope string    `gorm:"size:64;default:'ALL'" json:"data_scope"`
	Remark    string    `json:"remark"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (r *SysRole) TableName() string {
	return "sys_role"
}
