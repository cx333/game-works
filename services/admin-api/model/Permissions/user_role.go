package Permissions

import "github.com/cx333/game-works/services/admin-api/resource"

type PermUserRole struct {
	resource.Models
	UserID uint64 `json:"userID"`
	RoleID uint64 `json:"roleID"`
}

func (*PermUserRole) TableName() string {
	return "perm_user_role"
}
