package main

import "github.com/cx333/game-works/pkg/frame"

/**
 * @Author: wgl
 * @Description:
 * @File: main
 * @Version: 1.0.0
 * @Date: 2025/4/19 20:11
 */

func main() {
	frameTicker := frame.NewFrameLoop(20)
	frameTicker.Register(func() {

	})
	frameTicker.Start()
	// 创建房间
	manager := NewRoomManager()
	room, err := manager.CreateRoom("test-room01")
	if err != nil {
		return
	}
	room.editRoomPlayer(&Player{
		playerId: "user1",
		nickname: "玩家1",
		avatar:   "😊",
	})
}
