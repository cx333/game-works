package organize

import "github.com/cx333/game-works/services/admin-api/resource"

type User struct {
	Username string `gorm:"size:64;not null" json:"username"`
	Password string `gorm:"size:128;not null" json:"-"`
	RealName string `gorm:"size:64;" json:"realName"` // 实名
	Nickname string `gorm:"size:64;" json:"nickName"`
	Email    string `gorm:"size:64;" json:"email"`
	Phone    string `gorm:"size:64;" json:"phone"`
	Status   int    `gorm:"size:8;" json:"status"`
	Avatar   string `gorm:"size:256" json:"avatar"`
	Desc     string `gorm:"size:256" json:"desc"`
	resource.Models
}

func (u *User) TableName() string {
	return "user"
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
