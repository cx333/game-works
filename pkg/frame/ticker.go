package frame

import (
	"sync"
	"time"
)

/**
 * @Author: wgl
 * @Description: FrameLoop 实现
 * @File: ticker
 * @Version: 1.0.0
 * @Date: 2025/4/19 20:09
 */

// Loop FrameLoop 帧循环调度器结构体
type Loop struct {
	ticker    *time.Ticker      // 定时触发一帧
	interval  time.Duration     // 帧之间的间隔时间 用来构造 ticker
	callbacks []func(tick uint) // 保存每帧需要执行的回调函数 带tick
	mu        sync.RWMutex      // 保护 callbacks 并发安全
	stopChan  chan struct{}     // 停止信号通道
	tick      uint
}

// NewFrameLoop 创建一个新的调度器，fps 是帧率 （每秒执行几次）
func NewFrameLoop(fps int) *Loop {
	interval := time.Second / time.Duration(fps)
	return &Loop{
		interval: interval,
		stopChan: make(chan struct{}),
	}
}

// Register 注册一个每帧都会触发的函数
func (f *Loop) Register(cb func(tick uint)) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.callbacks = append(f.callbacks, cb)
}

// Start 启动帧循环
func (f *Loop) Start() {
	f.ticker = time.NewTicker(f.interval)
	go func() {
		for {
			select {
			case <-f.ticker.C:
				f.step()
			case <-f.stopChan:
				f.ticker.Stop()
			}
		}
	}()
}

// Stop 停止帧循环
func (f *Loop) Stop() {
	close(f.stopChan)
}

// step 每帧执行所有注册的回调
func (f *Loop) step() {
	f.mu.Lock()
	defer f.mu.Unlock()
	for _, cb := range f.callbacks {
		cb(f.tick)
	}
	f.tick++
}
