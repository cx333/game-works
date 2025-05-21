package main

import (
	"encoding/binary"
	"fmt"
	"github.com/cx333/game-works/pkg/logger"
	"github.com/cx333/game-works/pkg/proto"
	"github.com/cx333/game-works/pkg/shared"
	protouf "google.golang.org/protobuf/proto"
	"io"
	"net"
)

func client() net.Conn {
	conn, err := net.Dial("tcp", shared.TcpPort)
	if err != nil {
		fmt.Println("Dial error:", err)
		return nil
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			logger.Info("TCP client err ", err)
			return
		}
	}(conn)
	return conn
}

func protoMarshal(playerId, command string, payload []byte) (data []byte, err error) {
	// 创建 ClientMessage
	clientMessage := &proto.ClientMessage{
		PlayerId: playerId,
		Command:  command,
		Payload:  payload,
	}
	// 编码 ClientMessage 为二进制
	data, err = protouf.Marshal(clientMessage)
	if err != nil {
		return nil, err
	}
	return data, err
}

// 写入消息（添加长度前缀）
func writeMsg(conn net.Conn, msg []byte) error {
	header := make([]byte, 4)
	binary.BigEndian.PutUint32(header, uint32(len(msg)))

	if _, err := conn.Write(header); err != nil {
		return err
	}
	if _, err := conn.Write(msg); err != nil {
		return err
	}
	return nil
}

// 读取完整消息
func readMsg(conn net.Conn) ([]byte, error) {
	header := make([]byte, 4)
	if _, err := io.ReadFull(conn, header); err != nil {
		return nil, err
	}

	msgLen := binary.BigEndian.Uint32(header)
	if msgLen > 1024*1024 {
		return nil, fmt.Errorf("message too long: %d", msgLen)
	}

	msg := make([]byte, msgLen)
	if _, err := io.ReadFull(conn, msg); err != nil {
		return nil, err
	}

	return msg, nil
}
