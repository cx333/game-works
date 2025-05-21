package main

import (
	"github.com/cx333/game-works/pkg/logger"
	"github.com/cx333/game-works/pkg/model"
	"github.com/cx333/game-works/pkg/natsx"
	protocol "github.com/cx333/game-works/pkg/proto"
	"github.com/cx333/game-works/pkg/shared"
	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
	"time"
)

/**
 * @Author: wgl
 * @Description:
 * @File: main
 * @Version: 1.0.0
 * @Date: 2025/4/19 20:11
 */

func init() {
	logger.Init("room", logger.DebugLevel, "../logs/room")
	defer logger.Sync()
	config := natsx.NatsConfig{
		URL:            "nats://192.168.1.63:4222",
		Name:           "game-server",
		MaxReconnects:  -1, // 无限重连
		ReconnectWait:  2 * time.Second,
		ConnectTimeout: 5 * time.Second,
	}
	nc, err := natsx.New(config)
	if err != nil {
		logger.Warn("Failed to connect to NATS", err)
		return
	}
	shared.RoomNats = nc
}

func main() {
	err := shared.RoomNats.SubscribeWithRetry(natsx.RoomCreateTopic, func(msg *nats.Msg) {
		var roomMsg protocol.RoomMessage
		if err := proto.Unmarshal(msg.Data, &roomMsg); err != nil {
			logger.Error("unmarshal chat message error", zap.Error(err))
			return
		}

		switch roomMsg.Type {
		case shared.NewRoom:
			manager := NewRoomManager()
			room, err := manager.CreateRoom(roomMsg.RoomId, roomMsg.Password)
			if err != nil {
				return
			}
			shared.AllRoom.Store(roomMsg.RoomId, room)
			room.editRoomPlayer(&model.Player{
				PlayerId: roomMsg.PlayerId,
			})
		case shared.PushPlayer:
			if val, _ := shared.AllRoom.Load(roomMsg.RoomId); val != nil {
				room := val.(*Room)
				room.editRoomPlayer(&model.Player{
					PlayerId: roomMsg.PlayerId,
				})
			}
		case shared.ExitRoom:
			if val, _ := shared.AllRoom.Load(roomMsg.RoomId); val != nil {
				room := val.(*Room)
				room.deleteRoomPlayer(roomMsg.PlayerId)
			}
		default:
			return
		}
	})
	if err != nil {
		return
	}
}
