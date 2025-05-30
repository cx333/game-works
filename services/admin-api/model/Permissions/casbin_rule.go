package Permissions

type PermCasbinRule struct {
	ID    uint   `gorm:"primaryKey"`
	PType string `gorm:"size:100"`
	V0    string `gorm:"size:100"`
	V1    string `gorm:"size:100"`
	V2    string `gorm:"size:100"`
	V3    string `gorm:"size:100"`
	V4    string `gorm:"size:100"`
	V5    string `gorm:"size:100"`
}

func (*PermCasbinRule) TableName() string { return "perm_casbin_rule" }
