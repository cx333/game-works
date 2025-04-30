package transport

import (
	"bufio"
	"encoding/binary"
	"github.com/cx333/game-works/pkg/logger"
	"github.com/cx333/game-works/pkg/natsx"
	protocol "github.com/cx333/game-works/pkg/proto"
	"github.com/cx333/game-works/services/gateway/shared"
	"google.golang.org/protobuf/proto"
	"net"
	"os"
	"sync"
)

/**
 * @Author: wgl
 * @Description:
 * @File: tcp
 * @Version: 1.0.0
 * @Date: 2025/4/16 20:35
 */

var TcpConnMap = sync.Map{}

func StartTcpServer() {
	ln, err := net.Listen("tcp", shared.TcpPort)
	if err != nil {
		logger.Error("TCP server listen err:", err)
		os.Exit(1)
	}
	logger.Info("TCP服务监听启动 ", "port", shared.TcpPort)
	for {
		conn, err := ln.Accept()
		if err != nil {
			logger.Error("TCP server accept err:", err)
			continue
		}
		go handleConnectionTcp(conn)
	}
}

func handleConnectionTcp(conn net.Conn) {
	defer func() {
		err := conn.Close()
		if err != nil {
			logger.Warn("TCP connection close error:", err)
		}
	}()

	reader := bufio.NewReader(conn)
	for {
		// 读取消息长度
		lenBuf := make([]byte, 4)
		_, err := reader.Read(lenBuf)
		if err != nil {
			logger.Warn("TCP server read length error:", err)
			return
		}

		// 获取消息的实际长度
		length := binary.BigEndian.Uint32(lenBuf)
		if length == 0 {
			logger.Warn("Received message with length 0, ignoring.")
			continue
		}

		// 读取消息数据
		dataBuf := make([]byte, length)
		_, err = reader.Read(dataBuf)
		if err != nil {
			logger.Warn("TCP server read data error:", err)
			return
		}

		// 异步处理接收到的消息
		go func(data []byte) {
			// 反序列化 GatewayMessage
			var msg protocol.GatewayMessage
			if err := proto.Unmarshal(data, &msg); err != nil {
				logger.Warn("TCP server unmarshal error:", err)
				return
			}

			// 根据 cmd 路由到相应的服务
			serviceSubject := routeByCmd(msg.Cmd)
			if serviceSubject == "" {
				logger.Warn("No route found for cmd:", msg.Cmd)
				return
			}

			// 将 Payload 发布到 NATS
			err := shared.NatsConn.Publish(serviceSubject, msg.Payload)
			if err != nil {
				logger.Warn("TCP server publish error:", err)
				return
			}

			// 可选：打印或记录已成功发布的消息
			logger.Info("Published message to NATS. Cmd:", msg.Cmd, "Subject:", serviceSubject)
		}(dataBuf)
	}
}

func routeByCmd(cmd int32) string {
	switch cmd {
	case 1001:
		return "player.service"
	case 2001:
		return natsx.ChatSendTopic
	default:
		return "unknown"
	}
}
