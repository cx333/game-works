package room

import (
	"github.com/cx333/game-works/pkg/frame"
	"github.com/cx333/game-works/pkg/model"
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

// PlayerNum 房间玩家数量限制——暂定为2
const PlayerNum = 3

// Fps 暂定帧数为20
const Fps = 20

// Room 单房间结构体
type Room struct {
	// 互斥锁，保证并发安全
	mu sync.Mutex
	// 房间公共结构体
	rm *model.Room
	// 玩家映射表 [playerId]*model.Player
	players map[string]*model.Player
	// 当前帧(用于游戏状态同步)
	frameIndex int64
	// 帧循环控制器
	ticker *frame.Loop
	// 等待处理的输入
	pendingInputs []model.PlayerInput
}

type RoomImpl interface {
	editRoomPlayer(player *model.Player)
	deleteRoomPlayer(playerId string)
}

// editRoomPlayer 编辑房间玩家信息（添加、修改）
func (r *Room) editRoomPlayer(player *model.Player) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if len(r.players) < PlayerNum+1 {
		player.Hp = 100
		player.PosX = 0
		player.PosY = 0
		player.Action = model.Idle
		r.players[player.PlayerId] = player
	}
	// 当房间已满，转换状态为 房间准备就绪
	if len(r.players) == PlayerNum {
		r.rm.State = Ready
	}
}

// deleteRoomPlayer 删除房间玩家
func (r *Room) deleteRoomPlayer(playerId string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.players, playerId)
}
