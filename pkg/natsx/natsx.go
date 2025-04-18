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

var ErrNotConnected = fmt.Errorf("natsx: connection not initialized")

type NatsConn struct {
	Conn *nats.Conn
}

type NatsConnImpl interface {
	Publish(subject string, data []byte) error
	Subscribe(subject string, handler nats.MsgHandler) error
	Close()
}

// New 初始化连接并赋值给 Conn
func New(natsURL string) (*NatsConn, error) {
	conn, err := nats.Connect(natsURL,
		nats.MaxReconnects(10),
		nats.ReconnectWait(2*time.Second),
		nats.Name("game-server"),
	)
	if err != nil {
		log.Fatalf("Failed to connect to NATS: %v", err)
	}
	log.Println("NATS connected to", natsURL)
	return &NatsConn{Conn: conn}, nil
}

func (c *NatsConn) Publish(subject string, data []byte) error {
	if c.Conn == nil {
		return ErrNotConnected
	}
	return c.Conn.Publish(subject, data)
}

func (c *NatsConn) Subscribe(subject string, handler nats.MsgHandler) error {
	if c.Conn == nil {
		return ErrNotConnected
	}
	_, err := c.Conn.Subscribe(subject, handler)
	if err != nil {
		log.Println("NATS Subscribe error:", err)
	}
	return err
}

func (c *NatsConn) Close() {
	c.Conn.Close()
}
