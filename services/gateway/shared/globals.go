package shared

import "github.com/cx333/game-works/pkg/natsx"

/**
 * @Author: wgl
 * @Description:
 * @File: shared
 * @Version: 1.0.0
 * @Date: 2025/4/23 21:16
 */

var NatsConn = &natsx.NatsConn{}

const (
	TcpPort       = ":9000"
	UdpPort       = ":9001"
	WebSocketPort = ":9002"
	AdminPort     = ":9100"
	DebugPort     = ":9191"
)
