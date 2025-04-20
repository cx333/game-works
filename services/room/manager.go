package room

import (
	"fmt"
	"github.com/cx333/game-works/pkg/logger"
	"github.com/cx333/game-works/pkg/model"
	"sync"
)

/**
 * @Author: wgl
 * @Description: 管理所有房间
 * @File: manager
 * @Version: 1.0.0
 * @Date: 2025/4/19 20:16
 */

type RoomManager struct {
	mu    sync.Mutex
	rooms map[string]*Room
}

// NewRoomManager 创建房间 Manager 实例
func NewRoomManager() *RoomManager {
	return &RoomManager{
		rooms: make(map[string]*Room),
	}
}

// CreateRoom 创建房间（游戏中需要先创建房间再添加数据）
func (rm *RoomManager) CreateRoom(roomId string, passwd string) (*Room, error) {
	rm.mu.Lock()
	defer rm.mu.Unlock()
	if _, exists := rm.rooms[roomId]; exists {
		logger.Warn("创建失败，房间已存在")
		return nil, fmt.Errorf("创建失败，%s房间已存在", roomId)
	}
	// 创建新房间
	room := &Room{
		roomId:     roomId,
		state:      Wait,
		players:    make(map[string]*model.Player),
		passwd:     passwd,
		frameIndex: 0,
	}
	//room.StartFrameLoop()
	rm.rooms[roomId] = room
	return room, nil
}

// GetRoom 查询房间
func (rm *RoomManager) GetRoom(roomId string) (*Room, error) {
	rm.mu.Lock()
	defer rm.mu.Unlock()
	room, exists := rm.rooms[roomId]
	if !exists {
		log := fmt.Errorf("房间不存在")
		logger.Warn(log)
		return room, log
	}
	return room, nil
}

// RemoveRoom 删除房间
func (rm *RoomManager) RemoveRoom(roomId string) {
	rm.mu.Lock()
	defer rm.mu.Unlock()
	delete(rm.rooms, roomId)
}

// ForEachRoom 遍历房间执行逻辑
func (rm *RoomManager) ForEachRoom(f func(room *Room)) {
	rm.mu.Lock()
	defer rm.mu.Unlock()
	for _, room := range rm.rooms {
		f(room)
	}
}
