package organize

import (
	"github.com/cx333/game-works/services/admin-api/resource"
)

type Dept struct {
	ParentID uint64 `gorm:"default:0" json:"parent_id"`
	Name     string `gorm:"size:64;not null" json:"name"`
	Leader   string `gorm:"size:64" json:"leader"`
	Phone    string `gorm:"size:32" json:"phone"`
	Email    string `gorm:"size:128" json:"email"`
	Sort     int    `gorm:"default:0" json:"sort"`
	Status   int    `gorm:"default:1" json:"status"` // 0禁用, 1正常
	resource.Models
}

func (d *Dept) TableName() string {
	return "dept"
}
