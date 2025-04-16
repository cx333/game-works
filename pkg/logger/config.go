package logger

/**
 * @Author: wgl
 * @Description: 日志格式
 * @File: config
 * @Version: 1.0.0
 * @Date: 2025/4/16 22:37
 */

// LogLevel 日志等级
type LogLevel string

const (
	DebugLevel LogLevel = "debug"
	InfoLevel  LogLevel = "info"
	WarnLevel  LogLevel = "warn"
	ErrorLevel LogLevel = "error"
)

// SetLevel 动态设置日志等级
func SetLevel(level LogLevel) error {
	return atomicLevel.UnmarshalText([]byte(level))
}

// AllLevels 所有支持的等级（可用于前端下拉选项）
func AllLevels() []LogLevel {
	return []LogLevel{
		DebugLevel,
		InfoLevel,
		WarnLevel,
		ErrorLevel,
	}
}
