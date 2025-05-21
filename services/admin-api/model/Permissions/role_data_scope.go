package Permissions

import "github.com/cx333/game-works/services/admin-api/resource"

type PermRoleDataScope struct {
	resource.Models
	RoleID uint64 `json:"roleID"`
	DeptID uint64 `json:"deptID"`
}

func (*PermRoleDataScope) TableName() string {
	return "perm_role_data_scope"
}
