package shared

import (
	"github.com/cx333/game-works/pkg/natsx"
	"sync"
)

// room 服务

var RoomNats = &natsx.NatsConn{}

// 动作
const (
	NewRoom    = 1 // 创建房间
	PushPlayer = 2 // 进入房间
	ExitRoom   = 3 // 退出房间
)

// AllRoom 存储当前所有房间
var AllRoom sync.Map
