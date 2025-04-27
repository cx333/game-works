package transport

import (
	"bufio"
	"encoding/binary"
	"github.com/cx333/game-works/pkg/logger"
	protocol "github.com/cx333/game-works/pkg/proto"
	"github.com/cx333/game-works/services/gateway/shared"
	"google.golang.org/protobuf/proto"
	"net"
	"os"
)

/**
 * @Author: wgl
 * @Description:
 * @File: tcp
 * @Version: 1.0.0
 * @Date: 2025/4/16 20:35
 */

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
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			logger.Warn("TCP connection close err:", err)
			return
		}
	}(conn)
	reader := bufio.NewReader(conn)
	for {
		lenBuf := make([]byte, 4)
		_, err := reader.Read(lenBuf)
		if err != nil {
			logger.Warn("TCP server read err:", err)
			return
		}
		length := binary.BigEndian.Uint32(lenBuf)
		dataBuf := make([]byte, length)
		_, err = reader.Read(dataBuf)
		if err != nil {
			logger.Warn("TCP server read err:", err)
			return
		}
		go func() {
			var msg protocol.GatewayMessage
			if err := proto.Unmarshal(dataBuf, &msg); err != nil {
				logger.Warn("TCP server unmarshal err:", err)
				return
			}
			serviceSubject := routeByCmd(msg.Cmd)
			err := shared.NatsConn.Publish(serviceSubject, msg.Payload)
			if err != nil {
				logger.Warn("TCP server publish err:", err)
				return
			}
		}()
	}
}

func routeByCmd(cmd int32) string {
	switch cmd {
	case 1001:
		return "player.service"
	case 2001:
		return "chat.service"
	default:
		return "unknown"
	}
}
