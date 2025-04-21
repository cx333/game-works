package model

import "github.com/cx333/game-works/pkg/logger"

/**
 * @Author: wgl
 * @Description:
 * @File: player
 * @Version: 1.0.0
 * @Date: 2025/4/20 21:35
 */
const UnknownPlayer = "Unknown"

// 玩家动作类型常量
const (
	Idle    = "idle"    // 闲置
	Move    = "move"    // 移动
	Attack  = "attack"  // 攻击
	Defend  = "defend"  // 防御
	Jump    = "jump"    // 跳跃
	Crouch  = "crouch"  // 蹲下
	UseItem = "useItem" // 使用物品
	Skill   = "skill"   // 技能
	Emote   = "emote"   // 表情动作
	Quit    = "quit"    // 退出房间
)

var ActionToCode = map[string]int{
	Idle:    0,  // 闲置
	Move:    1,  // 移动
	Attack:  2,  // 攻击
	Defend:  3,  // 防御
	Jump:    4,  // 跳跃
	Crouch:  5,  // 蹲下
	UseItem: 6,  // 使用物品
	Skill:   7,  // 技能
	Emote:   9,  // 表情动作
	Quit:    10, // 退出房间
}

var CodeToAction = map[int]string{
	0:  "idle",
	1:  "move",
	2:  "attack",
	3:  "defend",
	4:  "jump",
	5:  "crouch",
	6:  "useItem",
	7:  "skill",
	9:  "emote",
	10: "quit",
}

// Player 玩家结构体
type Player struct {
	PlayerId string // 编号
	Nickname string // 昵称
	Avatar   string // 头像
	Hp       int    // 血量
	PosX     int    // 位置x
	PosY     int    // 位置y
	Action   string // 动作指令
	CampId   int    // 阵营
	RoleId   int    // 角色id
}

// PlayerInput 玩家操作结构体
type PlayerInput struct {
	PlayerId string                 // 哪个玩家发起的操作
	OpType   int                    // 操作类型：move、attack、skill 等
	Payload  map[string]interface{} // 附加数据，如目标坐标、技能ID等
	Frame    int64                  // 客户端发送时的帧号（可选，用于同步校验）
}

// GetActionName 获取动作名称
func GetActionName(code int) string {
	if name, ok := CodeToAction[code]; ok {
		return name
	}
	logger.Warn(code, "获取动作名称失败")
	return UnknownPlayer
}

// GetActionCode 获取动作代码
func GetActionCode(name string) int {
	if code, ok := ActionToCode[name]; ok {
		return code
	}
	logger.Warn(name, "获取动作编码失败")
	return -1
}
