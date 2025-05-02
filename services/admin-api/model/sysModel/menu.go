package sysModel

import "time"

type SysMenu struct {
	ID             int64     `gorm:"primaryKey;autoIncrement" json:"id"`             // 主键
	ParentID       int64     `gorm:"index;default:0" json:"parent_id"`               // 父级菜单 ID（顶级为 0）
	Title          string    `gorm:"type:varchar(100);not null" json:"title"`        // 菜单标题
	Path           string    `gorm:"type:varchar(255);not null" json:"path"`         // 路由路径
	Component      string    `gorm:"type:varchar(255)" json:"component"`             // 组件路径（字符串），前端动态映射
	Redirect       string    `gorm:"type:varchar(255)" json:"redirect,omitempty"`    // 重定向路径（可选）
	Name           string    `gorm:"type:varchar(100)" json:"name"`                  // 路由名称
	Icon           string    `gorm:"type:varchar(100)" json:"icon,omitempty"`        // 图标
	ActiveIcon     string    `gorm:"type:varchar(100)" json:"active_icon,omitempty"` // 激活状态图标
	ActivePath     string    `gorm:"type:varchar(255)" json:"active_path,omitempty"` // 激活父级菜单路径
	Order          int       `gorm:"default:0" json:"order"`                         // 菜单排序
	AffixTab       bool      `gorm:"default:false" json:"affix_tab"`                 // 是否固定标签页
	AffixTabOrder  int       `gorm:"default:0" json:"affix_tab_order"`               // 固定标签顺序
	KeepAlive      bool      `gorm:"default:false" json:"keep_alive"`                // 是否缓存
	HideInMenu     bool      `gorm:"default:false" json:"hide_in_menu"`              // 是否在菜单中隐藏
	HideInTab      bool      `gorm:"default:false" json:"hide_in_tab"`               // 是否在标签页隐藏
	HideChildren   bool      `gorm:"default:false" json:"hide_children_in_menu"`     // 子路由是否隐藏
	HideBreadcrumb bool      `gorm:"default:false" json:"hide_in_breadcrumb"`        // 是否在面包屑隐藏
	IgnoreAccess   bool      `gorm:"default:false" json:"ignore_access"`             // 是否忽略权限验证
	MenuVisible403 bool      `gorm:"default:false" json:"menu_visible_with_403"`     // 菜单可见但跳转到403
	OpenInNewWin   bool      `gorm:"default:false" json:"open_in_new_window"`        // 是否新窗口打开
	NoBasicLayout  bool      `gorm:"default:false" json:"no_basic_layout"`           // 是否不使用基础布局
	IframeSrc      string    `gorm:"type:varchar(255)" json:"iframe_src,omitempty"`  // iframe地址
	Link           string    `gorm:"type:varchar(255)" json:"link,omitempty"`        // 外链
	MaxOpenTab     int       `gorm:"default:-1" json:"max_num_of_open_tab"`          // 最大标签页数量
	Authority      string    `gorm:"type:varchar(255)" json:"authority"`             // 拥有访问权限的角色，用逗号分隔
	CreatedAt      time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (m *SysMenu) TableName() string {
	return "sys_menu"
}

type RouteMeta struct {
	Title           string   `json:"title"`
	Icon            string   `json:"icon,omitempty"`
	Order           int      `json:"order,omitempty"`
	AffixTab        bool     `json:"affixTab,omitempty"`
	Authority       []string `json:"authority,omitempty"`
	HideInMenu      bool     `json:"hideInMenu,omitempty"`
	KeepAlive       bool     `json:"keepAlive,omitempty"`
	OpenInNewWindow bool     `json:"openInNewWindow,omitempty"`
	NoBasicLayout   bool     `json:"noBasicLayout"` // 部分特殊页面如果不需要基础布局（页面顶部和侧边栏），可将noBasicLayout设置为true
}

type RouteRecordStringComponent struct {
	Name      string                       `json:"name"`
	Path      string                       `json:"path"`
	Component string                       `json:"component,omitempty"`
	Redirect  string                       `json:"redirect,omitempty"`
	Meta      *RouteMeta                   `json:"meta,omitempty"`
	Children  []RouteRecordStringComponent `json:"children,omitempty"`
}

func t() {
	_ = []RouteRecordStringComponent{
		{
			Name:     "Dashboard",
			Path:     "/",
			Redirect: "/analytics",
			Meta: &RouteMeta{
				Order: -1,
				Title: "page.dashboard.title",
			},
			Children: []RouteRecordStringComponent{
				{
					Name:      "Analytics",
					Path:      "/analytics",
					Component: "/dashboard/analytics/index", // 注意路径匹配 pageMap
					Meta: &RouteMeta{
						AffixTab: true,
						Title:    "page.dashboard.analytics",
					},
				},
				{
					Name:      "Workspace",
					Path:      "/workspace",
					Component: "/dashboard/workspace/index",
					Meta: &RouteMeta{
						Title: "page.dashboard.workspace",
					},
				},
			},
		},
		{
			Name:      "Test",
			Path:      "/test",
			Component: "/test/index",
			Meta: &RouteMeta{
				Title:         "page.test",
				NoBasicLayout: true, // 如果你支持这个字段，记得加在 struct 里
			},
		},
	}
}
