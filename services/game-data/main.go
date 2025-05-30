package game_data

import (
	"fmt"
	"github.com/cx333/game-works/pkg/frame"
)

/**
 * @Author: wgl
 * @Description:
 * @File: main
 * @Version: 1.0.0
 * @Date: 2025/4/19 20:05
 */

func InitRoom() {
	loop := frame.NewFrameLoop(20) // 20帧每秒
	loop.Register(Update)          // 注册 update 函数
	loop.Start()
}

func Update(tick uint) {
	// 每一帧要处理的逻辑：同步状态
	fmt.Println("当前是第", tick, "帧")
}
