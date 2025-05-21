package resource

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Models struct {
	ID        uint64         `gorm:"primaryKey;autoIncrement" json:"id"` // 主键
	CompanyID uint           `gorm:"not null" json:"companyID"`          // 所属公司ID
	DeptID    uint           `gorm:"not null" json:"deptID"`             // 所属部门ID
	CreatedBy uint64         `gorm:"not null" json:"createdBy"`
	UpdatedBy uint64         `gorm:"not null" json:"updatedBy"`
	DeletedBy *uint64        `gorm:"" json:"deletedBy"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

func (m *Models) BeforeCreate(tx *gorm.DB) (err error) {
	// 创建前逻辑：比如填充创建人（CreatedBy）和更新时间（UpdatedAt）
	// 假设从 tx.Statement.Context 中获取当前用户 ID
	if userID, ok := tx.Statement.Context.Value("userID").(uint64); ok {
		m.CreatedBy = userID
		m.UpdatedBy = userID
	} else {
		// 如果没有用户ID，使用默认值
		m.CreatedBy = 0
		m.UpdatedBy = 0
	}
	if companyID, ok := tx.Statement.Context.Value("companyID").(uint); ok {
		m.CompanyID = companyID
	} else {
		m.CompanyID = 0
	}
	if deptID, ok := tx.Statement.Context.Value("deptID").(uint); ok {
		m.DeptID = deptID
	} else {
		m.DeptID = 0
	}
	// 这里可以添加更多创建前的操作，如日志记录、验证等
	fmt.Println("BeforeCreate:", m.CreatedBy)
	return nil
}

func (m *Models) AfterCreate(tx *gorm.DB) (err error) {
	// 创建后逻辑：可以用来记录操作日志或触发其他操作
	fmt.Println("AfterCreate: 新数据创建成功，ID:", m.CreatedBy)
	// 如果需要通知其他系统或服务，可以在这里调用相关方法
	return nil
}

func (m *Models) BeforeUpdate(tx *gorm.DB) (err error) {
	// 更新前逻辑：比如填充更新人（UpdatedBy）
	if userID, ok := tx.Statement.Context.Value("userID").(uint64); ok {
		m.UpdatedBy = userID
	} else {
		m.UpdatedBy = 0
	}
	if companyID, ok := tx.Statement.Context.Value("companyID").(uint); ok {
		m.CompanyID = companyID
	} else {
		m.CompanyID = 0
	}
	if deptID, ok := tx.Statement.Context.Value("deptID").(uint); ok {
		m.DeptID = deptID
	} else {
		m.DeptID = 0
	}
	// 更新前的一些操作，类似创建时的操作
	fmt.Println("BeforeUpdate:", m.UpdatedBy)
	return nil
}

func (m *Models) AfterUpdate(tx *gorm.DB) (err error) {
	// 更新后逻辑：可以用来记录日志、通知等
	fmt.Println("AfterUpdate: 数据更新成功，更新人:", m.UpdatedBy)
	return nil
}

func (m *Models) BeforeDelete(tx *gorm.DB) (err error) {
	// 删除前逻辑：如果启用软删除，记录删除人（DeletedBy）
	if userID, ok := tx.Statement.Context.Value("userID").(uint64); ok {
		m.DeletedBy = &userID
	} else {
		m.DeletedBy = nil
	}
	if companyID, ok := tx.Statement.Context.Value("companyID").(uint); ok {
		m.CompanyID = companyID
	} else {
		m.CompanyID = 0
	}
	if deptID, ok := tx.Statement.Context.Value("deptID").(uint); ok {
		m.DeptID = deptID
	} else {
		m.DeptID = 0
	}
	fmt.Println("BeforeDelete:", m.DeletedBy)
	return nil
}

func (m *Models) AfterDelete(tx *gorm.DB) (err error) {
	// 删除后逻辑：执行删除后的操作，例如日志记录等
	fmt.Println("AfterDelete: 数据已删除，删除人:", m.DeletedBy)
	return nil
}
