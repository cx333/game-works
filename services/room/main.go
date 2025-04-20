package main

import (
	"github.com/cx333/game-works/pkg/logger"
	"github.com/cx333/game-works/pkg/model"
)

/**
 * @Author: wgl
 * @Description:
 * @File: main
 * @Version: 1.0.0
 * @Date: 2025/4/19 20:11
 */

func main() {
	logger.Init("room", logger.DebugLevel, "./logs")
	defer logger.Sync()
	// 创建房间
	manager := NewRoomManager()
	room, err := manager.CreateRoom("test-room01", "")
	if err != nil {
		return
	}
	// 添加房间玩家
	room.editRoomPlayer(&model.Player{
		PlayerId: "user1",
		Nickname: "玩家1",
		Avatar:   "😊",
	})
	room.editRoomPlayer(&model.Player{
		PlayerId: "user2",
		Nickname: "玩家2",
		Avatar:   "😊",
	})
	room.editRoomPlayer(&model.Player{
		PlayerId: "user3",
		Nickname: "玩家3",
		Avatar:   "😊",
	})
	room.StartGame()

	select {}
}
