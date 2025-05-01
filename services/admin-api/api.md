# 前端接口文档（apps/web-antd）

---

## 1. 认证相关接口（auth.ts）

### 登录
- 方法：`loginApi(data: { username: string; password: string })`
- 请求：`POST /sys/login`
- 参数：
  - `username`（string，必填）：用户名
  - `password`（string，必填）：密码
- 返回：
  ```json
  {
    "id": 1,
    "realName": "Admin",
    "roles": ["admin"],
    "username": "admin",
    "homePath": "/workspace",
    "accessToken": "string"
  }
  ```
  - 说明：返回用户基本信息和 accessToken

### 刷新 Token
- 方法：`refreshTokenApi()`
- 请求：`POST /auth/refresh`
- 参数：无（自动带上 cookie）
- 返回：
  ```json
  {
    "data": "string",   // 新的 accessToken
    "status": 0          // 状态码
  }
  ```
- 说明：带 `withCredentials: true`

### 退出登录
- 方法：`logoutApi()`
- 请求：`POST /auth/logout`
- 参数：无
- 返回：无数据（仅状态码）
- 说明：带 `withCredentials: true`

### 获取用户权限码
- 方法：`getAccessCodesApi()`
- 请求：`GET /auth/codes`
- 返回：
  ```json
  [
    "AC_100010", "AC_100020", ...
  ]
  ```
  - 说明：权限码根据用户角色动态返回，详见 mock-data.ts 中 MOCK_CODES。

---

## 2. 菜单相关接口（menu.ts）

### 获取用户所有菜单
- 方法：`getAllMenusApi()`
- 请求：`GET /menu/all`
- 返回：数组，每项结构如下（递归结构）：
  ```json
  {
    "component": "BasicLayout",
    "meta": {
      "order": -1,
      "title": "page.dashboard.title"
    },
    "name": "Dashboard",
    "path": "/",
    "redirect": "/analytics",
    "children": [
      {
        "name": "Analytics",
        "path": "/analytics",
        "component": "/dashboard/analytics/index",
        "meta": {
          "affixTab": true,
          "title": "page.dashboard.analytics"
        }
      }
      // ... 其它子菜单
    ]
  }
  ```
- 说明：菜单结构会根据用户角色动态变化，字段详见 RouteMeta 类型定义，mock 示例详见 mock-data.ts 中 MOCK_MENUS。

---

## 3. 用户相关接口（user.ts）

### 获取用户信息
- 方法：`getUserInfoApi()`
- 请求：`GET /user/info`
- 返回：
  ```json
  {
    "id": 1,
    "realName": "Admin",
    "roles": ["admin"],
    "username": "admin",
    "homePath": "/workspace"
  }
  ```
- 说明：字段含义详见 mock-data.ts UserInfo 接口。

---

## 4. 其它说明

- 所有接口请求通过 `requestClient` 或 `baseRequestClient` 发送，自动携带 token、语言等信息。
- 响应自动处理 code、data 字段，code=0 时返回 data，否则抛出异常。
- 支持 token 过期自动刷新、统一错误提示。
- 权限码、菜单、用户信息等均为 mock 数据，实际项目可对接真实后端。

---

## 5. 其它页面相关

- 忘记密码、注册等页面目前仅做表单演示，未实际发起 API 请求。如需扩展，请在 `/api` 目录下新增对应接口。
- 相关表单字段和校验规则可参考 `src/views/_core/authentication/` 下的 vue 文件。

---

如需详细的参数类型或返回结构说明，可查看 `@vben/types`、`mock-data.ts` 中的类型定义。如需补充其它业务接口，请告知！

---

## 6. 后端数据表设计（Go + Gorm + PostgreSQL 格式）

### 用户表（users）
```go
// 用户信息表
// TableName: users
// 说明：存储系统用户的基本信息

type User struct {
    ID        uint           `gorm:"primaryKey;autoIncrement"` // 用户ID，主键
    Username  string         `gorm:"uniqueIndex;size:64;not null"` // 用户名，唯一
    Password  string         `gorm:"size:128;not null"` // 密码（加密存储）
    RealName  string         `gorm:"size:64;not null"` // 真实姓名
    Avatar    string         `gorm:"size:256"` // 头像URL
    HomePath  string         `gorm:"size:128"` // 登录后首页路径
    Desc      string         `gorm:"size:256"` // 用户描述
    CreatedAt time.Time      // 创建时间
    UpdatedAt time.Time      // 更新时间
    Roles     []Role         `gorm:"many2many:user_roles;"` // 用户角色，多对多
}
```

### 角色表（roles）
```go
// 角色信息表
// TableName: roles
// 说明：定义系统中的角色

type Role struct {
    ID        uint      `gorm:"primaryKey;autoIncrement"` // 角色ID，主键
    Name      string    `gorm:"uniqueIndex;size:64;not null"` // 角色标识，唯一
    Label     string    `gorm:"size:64;not null"` // 角色名称（显示用）
    Desc      string    `gorm:"size:256"` // 角色描述
    CreatedAt time.Time // 创建时间
    UpdatedAt time.Time // 更新时间
    Users     []User    `gorm:"many2many:user_roles;"` // 拥有该角色的用户
    Codes     []Code    `gorm:"many2many:role_codes;"` // 角色拥有的权限码
}
```

### 权限码表（codes）
```go
// 权限码表
// TableName: codes
// 说明：定义系统中的权限码（如菜单、按钮权限）

type Code struct {
    ID        uint      `gorm:"primaryKey;autoIncrement"` // 权限码ID，主键
    Code      string    `gorm:"uniqueIndex;size:64;not null"` // 权限码，唯一
    Desc      string    `gorm:"size:256"` // 权限码描述
    CreatedAt time.Time // 创建时间
    UpdatedAt time.Time // 更新时间
    Roles     []Role    `gorm:"many2many:role_codes;"` // 拥有该权限码的角色
}
```

### 菜单表（menus）
```go
// 菜单表
// TableName: menus
// 说明：系统菜单，支持多级嵌套

type Menu struct {
    ID        uint      `gorm:"primaryKey;autoIncrement"` // 菜单ID，主键
    Name      string    `gorm:"size:64;not null"` // 路由名
    Path      string    `gorm:"size:128;not null"` // 路由路径
    Component string    `gorm:"size:128"` // 前端组件路径
    ParentID  *uint     `gorm:"index"` // 父级菜单ID，顶级为NULL
    Order     int       `gorm:"default:0"` // 排序
    Icon      string    `gorm:"size:64"` // 菜单图标
    Title     string    `gorm:"size:64;not null"` // 菜单标题
    Meta      datatypes.JSON `gorm:"type:jsonb"` // 其它元信息（如affixTab、badge等）
    CreatedAt time.Time // 创建时间
    UpdatedAt time.Time // 更新时间
    Roles     []Role    `gorm:"many2many:menu_roles;"` // 拥有该菜单权限的角色
}
```

### 关联表说明
- user_roles：用户与角色多对多关联表
- role_codes：角色与权限码多对多关联表
- menu_roles：菜单与角色多对多关联表

---

> 以上结构可根据实际业务扩展，如增加登录日志、操作日志、按钮权限等。字段类型和注释已适配 Gorm + PostgreSQL，支持多对多、嵌套、JSON 字段等。

---

## 7. 权限控制设计（Casbin RBAC）

### 1. Casbin 权限模型（model.conf 示例）
```ini
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

g = _, _

g2 = _, _

[role_definition]
g = _, _
g2 = _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act
```
- sub：请求的主体（如用户ID、角色名）
- obj：资源对象（如菜单路由、API路径、权限码）
- act：操作（如 read、write、delete、access 等）

### 2. Casbin 策略示例（policy.csv）
```csv
p, admin, /user/info, read
p, admin, /menu/all, read
p, admin, /sys/login, write
p, admin, /auth/codes, read
p, super, /user/info, read
p, super, /menu/all, read
p, super, /sys/login, write
p, super, /auth/codes, read
p, user, /user/info, read
p, user, /menu/all, read

# 用户与角色绑定
# g, 用户名, 角色名
g, jack, user
g, admin, admin
g, vben, super
```

### 3. 与表结构的对应关系
- 用户（users）与角色（roles）通过 user_roles 关联，映射到 Casbin 的 g 规则（g, 用户, 角色）。
- 角色（roles）与权限（如菜单、API、权限码）通过 policy.csv 的 p 规则进行授权。
- 可将 obj 设计为 API 路径、菜单路由、权限码等，act 可为 read、write、access 等操作。

### 4. 说明
- 可根据实际业务扩展 Casbin 的模型，如支持资源层级、数据权限等。
- Casbin 策略可存储在数据库（推荐使用 casbin 官方适配器），也可用文件存储。
- 前端权限码、菜单、API 路由等均可纳入 Casbin 控制，实现统一 RBAC 权限管理。

---
