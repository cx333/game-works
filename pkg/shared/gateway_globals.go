package shared

import "github.com/cx333/game-works/pkg/natsx"

// gateway 服务

var GatewayNats = &natsx.NatsConn{}

const (
	TcpPort       = ":9000"
	UdpPort       = ":9001"
	WebSocketPort = ":9002"
	AdminPort     = ":9100"
	DebugPort     = ":9191"
)

// 推送类型
const ()
