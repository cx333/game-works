package main

import (
	"fmt"
	"github.com/cx333/game-works/pkg/logger"
	"github.com/cx333/game-works/pkg/proto" // 确保路径正确
	"github.com/gorilla/websocket"
	protouf "google.golang.org/protobuf/proto"
	"log"
	"time"
)

func init() {
	logger.Init("player", logger.DebugLevel, "../logs/player")
	defer logger.Sync()
}

// 创建并发送 ClientMessage
func createAndSendMessage(ws *websocket.Conn, playerID, command, gameMode string) error {
	// 创建 MatchRequest
	matchRequest := &proto.MatchRequest{
		PlayerId: playerID,
		GameMode: gameMode,
	}

	// 编码 MatchRequest 为二进制
	matchPayload, err := protouf.Marshal(matchRequest)
	if err != nil {
		return fmt.Errorf("编码 MatchRequest 失败: %v", err)
	}

	// 创建 ClientMessage
	clientMessage := &proto.ClientMessage{
		PlayerId: playerID,
		Command:  command,
		Payload:  matchPayload,
	}

	// 编码 ClientMessage 为二进制
	data, err := protouf.Marshal(clientMessage)
	if err != nil {
		return fmt.Errorf("编码 ClientMessage 失败: %v", err)
	}

	// 发送二进制数据到 WebSocket 服务器
	err = ws.WriteMessage(websocket.BinaryMessage, data)
	if err != nil {
		return fmt.Errorf("发送 WebSocket 消息失败: %v", err)
	}

	log.Println("消息发送成功")
	return nil
}

func main() {
	// 连接到 WebSocket 服务器
	serverAddr := "ws://127.0.0.1:9001/ws" // 你网关的 WebSocket 地址
	ws, _, err := websocket.DefaultDialer.Dial(serverAddr, nil)
	if err != nil {
		log.Fatalf("连接 WebSocket 服务器失败: %v", err)
	}
	defer ws.Close()

	// 模拟发送的消息数据
	playerID := "1234"
	command := "gateway"
	gameMode := "hello"

	// 每隔 5 秒发送一次消息
	for {
		err := createAndSendMessage(ws, playerID, command, gameMode)
		if err != nil {
			log.Println("发送消息失败:", err)
		}

		// 等待 5 秒
		time.Sleep(5 * time.Second)
	}
}
