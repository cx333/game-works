package shared

import (
	"github.com/cx333/game-works/pkg/natsx"
	protocol "github.com/cx333/game-works/pkg/proto"
)

/**
 * @Author: wgl
 * @Description:
 * @File: globals
 * @Version: 1.0.0
 * @Date: 2025/4/24 16:33
 */

var NatsConn = &natsx.NatsConn{}

// ChatPrivateChan 私聊管道
var ChatPrivateChan = make(chan *protocol.ChatMessage, 100)

// ChatRoomChan 房间管道
var ChatRoomChan = make(chan *protocol.ChatMessage, 100)

// ChatPublicChan 广播管道
var ChatPublicChan = make(chan *protocol.ChatMessage, 100)
