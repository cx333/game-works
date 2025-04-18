package router

import (
	"github.com/cx333/game-works/pkg/logger"
	"github.com/cx333/game-works/pkg/proto"
	"github.com/gorilla/websocket"
	protobuf "google.golang.org/protobuf/proto"
)

/**
 * @Author: wgl
 * @Description: 路由
 * @File: router
 * @Version: 1.0.0
 * @Date: 2025/4/18 17:02
 */

// HandlerFunc 消息处理函数统一签名
// 每个 command 对应一个 handler， 负责解析 payload 和业务逻辑
type HandlerFunc func(conn *websocket.Conn, msg *proto.ClientMessage)

// 保存 command -> handle 的映射关系
var routers = map[string]HandlerFunc{}

// Register 注册一个新的消息处理器
func Register(command string, handler HandlerFunc) {
	routers[command] = handler
}

// HandleMessage 消息统一入口
// 负责反序列化 ClientMessage 分发到对应 handler
func HandleMessage(conn *websocket.Conn, raw []byte) {
	var msg proto.ClientMessage
	if err := protobuf.Unmarshal(raw, &msg); err != nil {
		logger.Warn("反序列化 ClientMessage 失败: ", err)
		return
	}

	// 尝试搜索对应类型
	handler, ok := routers[msg.Command]
	if !ok {
		logger.Warn("未注册的命令类型:", msg.Command)
		return
	}

	// 分发给对应 handle 处理
	handler(conn, &msg)
}
