package organize

import "github.com/cx333/game-works/services/admin-api/resource"

// 公司管理

type Company struct {
	resource.Models
	Name string `json:"name"`
	Code string `json:"code"`
}

func (*Company) TableName() string {
	return "company"
}
