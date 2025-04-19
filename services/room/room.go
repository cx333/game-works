package main

import (
	"github.com/cx333/game-works/pkg/frame"
	"sync"
)

/**
 * @Author: wgl
 * @Description: 单个房间结构和方法
 * @File: room
 * @Version: 1.0.0
 * @Date: 2025/4/19 20:16
 */

// 房间状态
const (
	// Wait 房间等待中
	Wait = 0
	// Ready 房间准备就绪
	Ready = 1
	// Running 游戏进行中
	Running = 2
	// Ended 本局已结束
	Ended = 3
)

// RoomPlayerNum 房间玩家数量限制——暂定为2
const RoomPlayerNum = 2

// Room 单房间结构体
type Room struct {
	mu      sync.Mutex
	roomId  string
	players map[string]*Player
	ticker  frame.FrameLoop
	state   int8
	passwd  string
}

type RoomImpl interface {
	editRoomPlayer(player *Player)
	deleteRoomPlayer(playerId string)
}

// editRoomPlayer 编辑房间玩家信息（添加、修改）
func (r *Room) editRoomPlayer(player *Player) {
	r.mu.Lock()
	defer r.mu.Unlock()
	// 当房间已满，转换状态为 房间准备就绪
	if len(r.players) == RoomPlayerNum {
		r.state = Ready
		return
	} else {
		r.players[player.playerId] = player
	}
}

// deleteRoomPlayer 删除房间玩家
func (r *Room) deleteRoomPlayer(playerId string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.players, playerId)
}
