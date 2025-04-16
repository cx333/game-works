package natsx

import "log"

/**
 * @Author: wgl
 * @Description:
 * @File: topic_registry
 * @Version: 1.0.0
 * @Date: 2025/4/16 22:06
 */

// TopicMeta 描述一个 Topic 的用途
type TopicMeta struct {
	Topic       string // 主题字符串
	Description string // 中文说明
}

// topicRegistry 当前服务主动注册的 Topic（用于订阅）
var topicRegistry = make([]TopicMeta, 0)

// RegisterTopic 注册一个订阅主题及说明
func RegisterTopic(topic string, description string) {
	topicRegistry = append(topicRegistry, TopicMeta{
		Topic:       topic,
		Description: description,
	})
}

// PrintRegisteredTopics 启动时打印所有注册的 Topic
func PrintRegisteredTopics(serviceName string) {
	log.Printf("🟢 [%s] 注册了以下 NATS Topic 订阅：\n", serviceName)
	for _, meta := range topicRegistry {
		log.Printf("   📥 %-20s | %s\n", meta.Topic, meta.Description)
	}
}
