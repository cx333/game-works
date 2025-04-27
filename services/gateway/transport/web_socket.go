package transport

import (
	"github.com/cx333/game-works/pkg/logger"
	"github.com/gorilla/websocket"
	"net/http"
)

/**
 * @Author: wgl
 * @Description:
 * @File: web_socket
 * @Version: 1.0.0
 * @Date: 2025/4/16 20:35
 */

var upgrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func StartWebSocketServer(addr string) {
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrade.Upgrade(w, r, nil)
		if err != nil {
			logger.Error("websocket连接升级失败:", err)
			return
		}
		logger.Info("websocket客户端已连接")
		go handleConnection(conn)
	})
	logger.Info("WebSocket 服务启动于: ", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		logger.Error("websocket 启动失败")
	}
}

func handleConnection(conn *websocket.Conn) {
	defer WsClose(conn)
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			logger.Warn("客户端断开连接：", err)
			return
		}
		logger.Debug("🫡收到消息：", string(msg))
		// 调用 router 路由消息
		//router.HandleMessage(conn, msg)
	}
}

// WsClose 断开连接
func WsClose(conn *websocket.Conn) {
	err := conn.Close()
	if err != nil {
		logger.Error("websocket close:", err)
		return
	}
	logger.Info("websocket 客户端已断开")
}
