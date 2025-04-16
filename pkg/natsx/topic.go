package natsx

import "fmt"

/**
 * @Author: wgl
 * @Description: 定义所有 NATS 使用的 Topic（支持手动常量和动态构造）
 * @File: topic
 * @Version: 1.0.0
 * @Date: 2025/4/16 21:58
 */

// TopicOf 构造标准化 topic，例如 TopicOf("room", "create") => "room.create"
func TopicOf(service string, action string) string {
	return fmt.Sprintf("%s.%s", service, action)
}

// ========= 以下是推荐使用的常量定义 =========

const (
	// ========= 匹配服务 =========
	MatchRequestTopic = "match.request" // 玩家请求匹配
	MatchResultTopic  = "match.result"  // 匹配成功后返回匹配结果

	// ========= 聊天服务 =========
	ChatSendTopic      = "chat.send"      // 客户端发送聊天消息
	ChatBroadcastTopic = "chat.broadcast" // 服务端广播聊天消息

	// ========= 房间服务 =========
	RoomCreateTopic = "room.create" // 创建房间
	RoomCloseTopic  = "room.close"  // 房间关闭

	// ========= 玩家状态服务 =========
	PlayerLoginTopic  = "player.login"  // 玩家登录
	PlayerLogoutTopic = "player.logout" // 玩家登出
	PlayerUpdateTopic = "player.update" // 玩家状态更新

	// ========= 监控服务 =========
	MonitorStatsTopic = "monitor.stats" // 服务状态上报

	// ========= 游戏数据服务 =========
	GameRechargeTopic = "game.recharge" // 充值记录
	GameGiftTopic     = "game.gift"     // 礼包发放

	// ========= 世界服务 =========
	WorldEventTopic = "world.event" // 活动或全服公告广播

	// ========= 鉴权服务 =========
	AuthLoginTopic  = "auth.login"  // 登录请求
	AuthVerifyTopic = "auth.verify" // 验证 token 合法性
)

var TopicDescriptions = map[string]string{
	MatchRequestTopic:  "玩家请求匹配",
	ChatBroadcastTopic: "聊天广播消息",
	// ...
}
