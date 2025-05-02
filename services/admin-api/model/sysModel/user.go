package sysModel

import "time"

type SysUser struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Username  string    `gorm:"uniqueIndex;size:64;not null" json:"username"`
	Password  string    `gorm:"size:128;not null" json:"-"`
	RealName  string    `gorm:"size:64;not null" json:"realName"` // 实名
	Nickname  string    `gorm:"size:64;" json:"nickName"`
	Email     string    `gorm:"size:64;" json:"email"`
	Phone     string    `gorm:"size:64;" json:"phone"`
	Status    int       `gorm:"size:8;" json:"status"`
	Avatar    string    `gorm:"size:256" json:"avatar"`
	HomePath  string    `gorm:"size:128" json:"homePath"`
	DeptID    int       `gorm:"size:8;" json:"deptID"`
	PostID    int       `gorm:"size:8;" json:"postID"`
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
	Desc     string `json:"desc"`
}

type UserInfoRes struct {
	Avatar   string   `json:"avatar"`
	RealName string   ` json:"realName"`
	Roles    []string ` json:"roles"`
	Id       string   `json:"userId"`
	Username string   `json:"username"`
}
