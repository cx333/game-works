package model

/**
 * @Author: wgl
 * @Description:
 * @File: player
 * @Version: 1.0.0
 * @Date: 2025/4/20 21:35
 */

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

var actionToCode = map[string]int{
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

var codeToAction = map[int]string{
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
