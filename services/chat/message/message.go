package message

import (
	"github.com/cx333/game-works/pkg/logger"
	"github.com/cx333/game-works/pkg/natsx"
	protocol "github.com/cx333/game-works/pkg/proto"
	"github.com/cx333/game-works/services/chat/shared"
	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
	"log"
)

/**
 * @Author: wgl
 * @Description:
 * @File: message
 * @Version: 1.0.0
 * @Date: 2025/4/28 16:33
 */

func SubMessage() {
	go ConsumptionChatMsgChan()
	err := shared.NatsConn.SubscribeWithRetry(natsx.ChatSendTopic, func(msg *nats.Msg) {
		var chatMsg protocol.ChatMessage
		if err := proto.Unmarshal(msg.Data, &chatMsg); err != nil {
			logger.Error("unmarshal chat message error", zap.Error(err))
			return
		}
		switch chatMsg.Channel {
		case 1:
			shared.ChatPrivateChan <- &chatMsg
		case 2:
			shared.ChatRoomChan <- &chatMsg
		case 3:
			shared.ChatPublicChan <- &chatMsg
		}
	})
	if err != nil {
		logger.Warn("sub nats ChatSendTopic error", err)
		return
	}
}

func ConsumptionChatMsgChan() {
	defer logger.Info("chat message consumer stopped")
	for {
		select {
		case msg, ok := <-shared.ChatPrivateChan:
			if !ok {
				log.Println("channel closed")
				return
			}
			chatPrivate(msg)
		case msg, ok := <-shared.ChatRoomChan:
			if !ok {
				log.Println("channel closed")
				return
			}
			chatRoom(msg) // 需要实现这个处理函数
		case msg, ok := <-shared.ChatPublicChan:
			if !ok {
				log.Println("channel closed")
				return
			}
			chatPublic(msg) // 需要实现这个处理函数
			//case <-time.After(5 * time.Second): // 超时处理
			//	log.Println("chat message timeout")
		}
	}
}

// 私聊
func chatPrivate(msg *protocol.ChatMessage) {

}

func chatRoom(msg *protocol.ChatMessage) {

}

// 广播消息
func chatPublic(msg *protocol.ChatMessage) {
	gatewayMsg := &protocol.GatewayMessage{
		Cmd:     2001,
		Payload: mustMarshal(msg),
	}
	// 发送给gateway，gateway收到后广播给所有在线客户端
	err := shared.NatsConn.Publish(natsx.ChatPublicTopic, mustMarshal(gatewayMsg))
	if err != nil {
		logger.Error("failed to publish broadcast message", zap.Error(err))
	}
}

func mustMarshal(pb proto.Message) []byte {
	data, err := proto.Marshal(pb)
	if err != nil {
		panic(err) // 这里直接panic掉，说明你proto定义有问题
	}
	return data
}

func SendMessage() {

}
