package sysModel

import "time"

type SysUser struct {
	ID        uint64    `gorm:"primaryKey" json:"id"`
	Username  string    `gorm:"uniqueIndex;size:64;not null" json:"username"`
	Password  string    `gorm:"size:128;not null" json:"-"`
	Nickname  string    `gorm:"size:64" json:"nickname"`
	Email     string    `gorm:"size:128" json:"email"`
	Phone     string    `gorm:"size:32" json:"phone"`
	Status    int       `gorm:"default:1" json:"status"` // 0禁用, 1正常
	DeptID    uint      `json:"dept_id"`                 // 所属部门
	PostID    uint      `json:"post_id"`                 // 岗位
	Avatar    string    `json:"avatar"`
	Remark    string    `json:"remark"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *SysUser) TableName() string {
	return "sys_user"
}

type Auth struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthRegister struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	DeptId   uint   `json:"dept_id"`
	PostId   uint   `json:"post_id"`
	Avatar   string `json:"avatar"`
	Remark   string `json:"remark"`
}
