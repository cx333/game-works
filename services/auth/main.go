package main

import (
	protocol "github.com/cx333/game-works/pkg/proto"
	"google.golang.org/protobuf/proto"
	"log"
	"net"
)

func main() {
	// 连接到网关服务器
	conn, err := net.Dial("tcp", "127.0.0.1:9000")
	if err != nil {
		log.Fatalf("连接网关失败: %v", err)
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			return
		}
	}(conn)

	// 构造要发送的业务消息（比如聊天）
	chatMsg := &protocol.ChatMessage{
		From:      "user01",
		To:        "user02", // 私聊频道
		RoomId:    "room01", // 房间频道
		Channel:   1,        // 频道类型：1 世界频道
		Content:   "你好，世界！",
		Timestamp: 0, // 让服务器来填
	}

	// 先把业务消息序列化
	payload, err := proto.Marshal(chatMsg)
	if err != nil {
		log.Fatalf("序列化业务消息失败: %v", err)
	}

	// 再构造 GatewayMessage（带 cmd）
	msg := &protocol.GatewayMessage{
		Cmd:     2001, // 假设 2001 是聊天
		Payload: payload,
	}

	// 序列化 GatewayMessage
	data, err := proto.Marshal(msg)
	if err != nil {
		log.Fatalf("序列化 GatewayMessage 失败: %v", err)
	}

	//
	size := len(data)
	sizeBuf := []byte{byte(size >> 24), byte(size >> 16), byte(size >> 8), byte(size)}
	_, err = conn.Write(sizeBuf)
	if err != nil {
		log.Fatalf("发送消息大小失败: %v", err)
	}
	_, err = conn.Write(data)
	if err != nil {
		log.Fatalf("发送消息内容失败: %v", err)
	}

	log.Println("消息发送成功！")
}
