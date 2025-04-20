package main

import (
	"fmt"
	"github.com/cx333/game-works/pkg/frame"
	"github.com/cx333/game-works/pkg/logger"
)

/**
 * @Author: wgl
 * @Description: 每个房间的逻辑帧调用
 * @File: frame
 * @Version: 1.0.0
 * @Date: 2025/4/19 20:18
 */

// StartGame 开始游戏
func (r *Room) StartGame() {
	if r.rm.State == Ready {
		r.rm.State = Running
		r.StartFrameLoop()
		log := fmt.Sprintf("房间 %s 开始游戏", r.rm.RoomId)
		logger.Info(log)
		return
	}
	logger.Warn("房间", r.rm.RoomId, "未准备就绪")
	return
}

// StartFrameLoop 启动房间帧循环
func (r *Room) StartFrameLoop() {
	r.ticker = frame.NewFrameLoop(Fps)
	r.ticker.Register(func(tick uint) {
		r.UpdateFrame()
	})
	r.ticker.Start()
}

// StopFrameLoop 停止房间帧循环
func (r *Room) StopFrameLoop() {
	if r.ticker != nil {
		r.ticker.Stop()
	}
}

func (r *Room) UpdateFrame() {
	r.mu.Lock()
	defer r.mu.Unlock()
	// 帧编号递增
	r.frameIndex++
	fmt.Printf("Room %s - Frame #%d: 玩家数: %d\n", r.rm.RoomId, r.frameIndex, len(r.players))
	// 逻辑
	for _, player := range r.players {
		fmt.Println("处理玩家", player.PlayerId)
	}
}
