package sysModel

import "time"

type SysUser struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Username  string    `gorm:"uniqueIndex;size:64;not null" json:"username"`
	Password  string    `gorm:"size:128;not null" json:"-"`
	RealName  string    `gorm:"size:64;not null" json:"realName"`
	Avatar    string    `gorm:"size:256" json:"avatar"`
	HomePath  string    `gorm:"size:128" json:"homePath"`
	Desc      string    `gorm:"size:256" json:"desc"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Roles     []SysRole `gorm:"many2many:user_roles;" json:"roles"`
}

func (u *SysUser) TableName() string {
	return "sys_user"
}

type Auth struct {
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
