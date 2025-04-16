package natsx

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"time"
)

/**
 * @Author: wgl
 * @Description:
 * @File: natsx
 * @Version: 1.0.0
 * @Date: 2025/4/16 20:50
 */

var Conn *nats.Conn

// Init 初始化连接并赋值给 Conn
func Init(natsURL string) {
	var err error
	Conn, err = nats.Connect(natsURL,
		nats.MaxReconnects(10),
		nats.ReconnectWait(2*time.Second),
		nats.Name("game-server"),
	)
	if err != nil {
		log.Fatalf("Failed to connect to NATS: %v", err)
	}
	log.Println("NATS connected to", natsURL)
}

func Subscribe(subject string, handler nats.MsgHandler) error {
	if Conn == nil {
		return ErrNotConnected
	}
	_, err := Conn.Subscribe(subject, handler)
	if err != nil {
		log.Println("NATS Subscribe error:", err)
	}
	return err
}

func Publish(subject string, data []byte) error {
	if Conn == nil {
		return ErrNotConnected
	}
	return Conn.Publish(subject, data)
}

var ErrNotConnected = fmt.Errorf("natsx: connection not initialized")
