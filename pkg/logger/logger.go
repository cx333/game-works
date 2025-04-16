package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

/**
 * @Author: wgl
 * @Description: 统一日志系统
 * @File: logger
 * @Version: 1.0.0
 * @Date: 2025/4/16 22:11
 */

var (
	log         *zap.SugaredLogger
	atomicLevel zap.AtomicLevel
)

func Init(serviceName string, level string, logDir string) {
	// 日志等级控制器
	atomicLevel = zap.NewAtomicLevel()
	_ = atomicLevel.UnmarshalText([]byte(level)) // 设置初始等级

	encoder := zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		TimeKey:       "ts",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "caller",
		MessageKey:    "msg",
		StacktraceKey: "stacktrace",
		EncodeTime:    zapcore.ISO8601TimeEncoder,
		EncodeLevel:   zapcore.CapitalColorLevelEncoder,
		EncodeCaller:  zapcore.ShortCallerEncoder,
	})

	// 文件输出（支持按天切分）
	fileWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   logDir + "/" + serviceName + ".log",
		MaxSize:    100, // MB
		MaxBackups: 7,
		MaxAge:     7,    // days
		Compress:   true, // gzip
	})

	// 控制台输出
	consoleWriter := zapcore.Lock(os.Stdout)

	core := zapcore.NewTee(
		zapcore.NewCore(encoder, fileWriter, atomicLevel),
		zapcore.NewCore(encoder, consoleWriter, atomicLevel),
	)

	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	log = logger.Sugar().With("module", serviceName)
}

// SetLevel 动态设置日志等级
func SetLevel(level string) error {
	return atomicLevel.UnmarshalText([]byte(level))
}

// Info 输出日志方法
func Info(args ...any)  { log.Infow("", "msg", args) }
func Error(args ...any) { log.Errorw("", "msg", args) }
func Debug(args ...any) { log.Debugw("", "msg", args) }
func Warn(args ...any)  { log.Warnw("", "msg", args) }

// Sync 强制刷盘（在退出时调用）
func Sync() {
	log.Sync()
}
