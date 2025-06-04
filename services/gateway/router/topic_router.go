package router

import (
	"github.com/cx333/game-works/pkg/natsx"
	"github.com/cx333/game-works/pkg/shared"
	"github.com/nats-io/nats.go"
	"net"
)

/**
 * @Author: wgl
 * @Description: 监听topic转发
 * @File: topic_router
 * @Version: 1.0.0
 * @Date: 2025/4/28 20:45
 */

func HandleAllMessage() {
	err := shared.GatewayNats.SubscribeWithRetry("chat.>", handleChatMessage)
	if err != nil {
		return
	}
}

var PublicChan chan map[net.Conn][]byte = make(chan map[net.Conn][]byte, 100)

func handleChatMessage(msg *nats.Msg) {
	switch msg.Subject {
	case "chat.send":
		return
	case natsx.ChatPrivateTopic:

	case natsx.ChatRoomTopic:

	// TODO: The following broadcast logic for ChatPublicTopic is incomplete.
	// It requires a consumer goroutine for PublicChan to process and send messages to clients.
	// The message framing (e.g., length-prefixing) also needs to be implemented for these broadcasts.
	case natsx.ChatPublicTopic:
		//go func() {
		//	transport.TcpConnMap.Range(func(key, value any) bool {
		//		if conn := value.(net.Conn); conn != nil {
		//			pubMap := make(map[net.Conn][]byte, 1)
		//			pubMap[conn] = msg.Data
		//			PublicChan <- pubMap
		//		}
		//		return true
		//	})
		//}()
	}
}
