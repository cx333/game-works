package model

/**
 * @Author: wgl
 * @Description:
 * @File: room
 * @Version: 1.0.0
 * @Date: 2025/4/20 21:35
 */

// Room 单房间结构体
type Room struct {
	// 房间唯一ID
	RoomId string
	// 房间状态(0:等待中 1:进行中 2:已结束)
	State int8
	// 房间密码(可选)
	Passwd string
	// 玩家数量
	PlayerNum int
	// 阵营
	Camps map[int]string
}
