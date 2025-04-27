package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

var (
	log         *zap.SugaredLogger
	atomicLevel zap.AtomicLevel
)

func Init(serviceName string, level LogLevel, logDir string) {
	atomicLevel = zap.NewAtomicLevel()
	_ = atomicLevel.UnmarshalText([]byte(level)) // 初始等级
	// 构造日志目录
	now := time.Now()
	dateStr := now.Format("2006-1-2:15") // 形如 2025-4-24:16
	//fullDir := logDir + "/" + serviceName
	if err := os.MkdirAll(logDir, 0755); err != nil {
		panic("无法创建日志目录: " + logDir)
	}
	filePath := logDir + "/" + dateStr + ".log"

	// 控制台格式：更适合阅读
	consoleEncoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		TimeKey:        "T",
		LevelKey:       "L",
		NameKey:        "N",
		CallerKey:      "C",
		MessageKey:     "M",
		StacktraceKey:  "S",
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
	})

	// 文件格式：结构化 JSON
	fileEncoder := zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
	})

	logFile := zapcore.AddSync(&lumberjack.Logger{
		Filename:   filePath,
		MaxSize:    100, // MB
		MaxBackups: 7,
		MaxAge:     7,
		Compress:   true,
	})

	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, logFile, atomicLevel),
		zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), atomicLevel),
	)

	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	log = logger.Sugar().With("service", serviceName)
}

// 基础方法（支持结构化参数）
func Info(msg string, args ...any)  { log.Infow(msg, args...) }
func Error(msg string, args ...any) { log.Errorw(msg, args...) }
func Debug(msg string, args ...any) { log.Debugw(msg, args...) }
func Warn(msg string, args ...any)  { log.Warnw(msg, args...) }

// 可选：格式化风格方法
func Infof(format string, args ...any)  { log.Infof(format, args...) }
func Errorf(format string, args ...any) { log.Errorf(format, args...) }
func Debugf(format string, args ...any) { log.Debugf(format, args...) }
func Warnf(format string, args ...any)  { log.Warnf(format, args...) }

func Sync() {
	_ = log.Sync()
}
