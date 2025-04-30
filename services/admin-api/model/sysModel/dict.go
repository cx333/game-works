package sysModel

import "time"

type SysDictType struct {
	ID        uint64    `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"size:64;not null" json:"name"`
	Type      string    `gorm:"size:64;not null;uniqueIndex" json:"type"`
	Status    int       `gorm:"default:1" json:"status"`
	Remark    string    `json:"remark"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type SysDictData struct {
	ID        uint64    `gorm:"primaryKey" json:"id"`
	TypeID    uint64    `json:"type_id"`
	Label     string    `gorm:"size:64;not null" json:"label"`
	Value     string    `gorm:"size:128;not null" json:"value"`
	Sort      int       `gorm:"default:0" json:"sort"`
	Status    int       `gorm:"default:1" json:"status"`
	Remark    string    `json:"remark"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (d *SysDictType) TableName() string {
	return "sys_dict_type"
}

func (d *SysDictData) TableName() string {
	return "sys_dict_data"
}
