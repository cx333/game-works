package main

import (
	"encoding/binary"
	"fmt"
	"github.com/cx333/game-works/pkg/logger"
	"github.com/cx333/game-works/pkg/proto" // 确保路径正确
	"github.com/gorilla/websocket"
	protouf "google.golang.org/protobuf/proto"
	"log"
	"net"
	"net/http"
	"time"
)

func init() {
	logger.Init("player", logger.DebugLevel, "../logs/player")
	defer logger.Sync()
}

// -------------------------------- 模拟客户端测试 -------------------------------------

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

func tcpCli(playerID, command, gameMode string) {
	// 连接到 TCP 服务器
	serverAddr := "127.0.0.1:9001" // 替换为你的 TCP 服务器地址和端口
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		log.Fatalf("连接 TCP 服务器失败: %v", err)
	}
	defer conn.Close()
	// 创建 MatchRequest
	//matchRequest := &proto.MatchRequest{
	//	PlayerId: playerID,
	//	GameMode: gameMode,
	//}
	// 编码 MatchRequest 为二进制
	//matchPayload, err := protouf.Marshal(matchRequest)
	// 构造包含 playerid 的消息
	clientMessage := &proto.ClientMessage{
		PlayerId: playerID,
		Command:  "0",
		Payload:  nil,
	}
	data, err := protouf.Marshal(clientMessage)
	if err != nil {
		log.Fatalf("序列化 playerid 消息失败: %v", err)
	}

	// 构造消息长度前缀
	length := uint32(len(data))
	lengthBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(lengthBytes, length)

	// 发送消息长度和消息内容
	_, err = conn.Write(lengthBytes)
	if err != nil {
		log.Fatalf("发送 playerid 消息长度失败: %v", err)
	}
	_, err = conn.Write(data)
	if err != nil {
		log.Fatalf("发送 playerid 消息内容失败: %v", err)
	}

	// 构造一条普通消息
	message := proto.GatewayMessage{
		Cmd:     "your_command",         // 替换为实际的命令
		Payload: []byte("your_payload"), // 替换为实际的负载
	}
	messageBytes, err := protouf.Marshal(&message)
	if err != nil {
		log.Fatalf("序列化消息失败: %v", err)
	}

	// 构造消息长度前缀
	length = uint32(len(messageBytes))
	lengthBytes = make([]byte, 4)
	binary.BigEndian.PutUint32(lengthBytes, length)

	// 发送消息长度和消息内容
	_, err = conn.Write(lengthBytes)
	if err != nil {
		log.Fatalf("发送消息长度失败: %v", err)
	}
	_, err = conn.Write(messageBytes)
	if err != nil {
		log.Fatalf("发送消息内容失败: %v", err)
	}

	// 等待服务器响应（可选）
	time.Sleep(5 * time.Second)
}

func main() {
	// 连接到 WebSocket 服务器
	serverAddr := "ws://127.0.0.1:9001/ws" // 你网关的 WebSocket 地址
	header := http.Header{}
	header.Add("player_id", "00101")
	ws, _, err := websocket.DefaultDialer.Dial(serverAddr, header)
	if err != nil {
		log.Fatalf("连接 WebSocket 服务器失败: %v", err)
	}
	defer func(ws *websocket.Conn) {
		err := ws.Close()
		if err != nil {
			log.Fatalf("defer WebSocket 服务器失败: %v", err)
		}
	}(ws)

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
