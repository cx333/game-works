package router

import (
	"github.com/cx333/game-works/pkg/natsx"
	"github.com/cx333/game-works/services/gateway/shared"
	"github.com/nats-io/nats.go"
)

/**
 * @Author: wgl
 * @Description: 监听topic转发
 * @File: topic_router
 * @Version: 1.0.0
 * @Date: 2025/4/28 20:45
 */

func HandleAllMessage() {
	err := shared.NatsConn.SubscribeWithRetry("chat.>", handleChatMessage)
	if err != nil {
		return
	}
}

func handleChatMessage(msg *nats.Msg) {
	switch msg.Subject {
	case "chat.send":
		return
	case natsx.ChatPrivateTopic:
	// 私聊
	case natsx.ChatRoomTopic:
	// 房间
	case natsx.ChatPublicTopic:
		// 广播

	}
}
