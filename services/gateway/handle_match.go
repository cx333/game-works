package main

import (
	"github.com/cx333/game-works/pkg/logger"
	"github.com/cx333/game-works/pkg/natsx"
	"github.com/cx333/game-works/pkg/proto"
	"github.com/gorilla/websocket"
)

/**
 * @Author: wgl
 * @Description: 注册服务、处理匹配消息
 * @File: handle_match
 * @Version: 1.0.0
 * @Date: 2025/4/18 20:42
 */
//
//func init() {
//	router.Register("match.request", matchRequestHandler)
//}

func matchRequestHandler(conn *websocket.Conn, msg *proto.ClientMessage) {
	logger.Debug("收到匹配请求，转发中...")

	// 把 payload 作为真正的消息发给 nats
	err := NatsConn.Publish(natsx.MatchRequestTopic, msg.Payload)
	if err != nil {
		logger.Error("转发匹配请求失败", err)
		return
	}
	logger.Debug("匹配请求转发成功")
}
