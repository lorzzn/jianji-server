package utils

import (
	"fmt"
	"jianji-server/config"
	"log"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger 全局日志指针
var Logger *zap.Logger

// 全局变量
var LoggerEncoder zapcore.Encoder
var LoggerFileSyncer zapcore.WriteSyncer
var LoggerConsoleSyncer zapcore.WriteSyncer
var LoggerWriterSyncer zapcore.WriteSyncer
var LoggerLevelEnabler zapcore.LevelEnabler

func SetupLogger() {
	// 生成日志文件目录
	if ok, _ := PathExists(config.Zap.Directory); !ok {
		log.Printf("创建日志目录：%v\n", config.Zap.Directory)
		_ = os.Mkdir(config.Zap.Directory, os.ModePerm)
	}

	//初始化全局变量
	LoggerEncoder = CreateLoggerEncoder()
	LoggerFileSyncer = CreateFileSyncer()
	LoggerConsoleSyncer = CreateConsoleSyncer()
	LoggerWriterSyncer = CreateLoggerWriterSyncer()
	LoggerLevelEnabler = CreateLoggerLevelPriority()
	//创建zap core
	core := zapcore.NewCore(LoggerEncoder, LoggerWriterSyncer, LoggerLevelEnabler)
	Logger = zap.New(core)

	if config.Zap.ShowLine {
		// 获取 调用的文件, 函数名称, 行号
		Logger = Logger.WithOptions(zap.AddCaller())
	}

	log.Println("Zap Logger 初始化成功")
}

// 编码器: 如何写入日志
func CreateLoggerEncoder() zapcore.Encoder {
	// 参考: zap.NewProductionEncoderConfig()
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     customTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder, // ?
	}

	if config.Zap.Format == "json" {
		return zapcore.NewConsoleEncoder(encoderConfig)
	}
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// 日志输出路径: 文件、控制台、双向输出
func CreateFileSyncer() zapcore.WriteSyncer {
	file, _ := os.Create(fmt.Sprintf("%s/server-%d.log", config.Zap.Directory, time.Now().UnixMilli()))
	return zapcore.AddSync(file)
}

func CreateConsoleSyncer() zapcore.WriteSyncer {
	return zapcore.AddSync(os.Stdout)
}

func CreateLoggerWriterSyncer() zapcore.WriteSyncer {
	// 双向输出
	if config.Zap.LogInConsole {
		return zapcore.NewMultiWriteSyncer(LoggerFileSyncer, LoggerConsoleSyncer)
	}

	// 输出到文件
	return CreateFileSyncer()
}

// 获取日志输出级别
func CreateLoggerLevelPriority() zapcore.LevelEnabler {
	switch config.Zap.Level {
	case "debug", "Debug":
		return zap.DebugLevel
	case "info", "Info":
		return zap.InfoLevel
	case "warn", "Warn":
		return zap.WarnLevel
	case "error", "Error":
		return zap.ErrorLevel
	case "dpanic", "DPanic":
		return zap.DPanicLevel
	case "panic", "Panic":
		return zap.PanicLevel
	case "fatal", "Fatal":
		return zap.FatalLevel
	}
	return zap.InfoLevel
}

// 自定义日志输出时间格式
func customTimeEncoder(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
	encoder.AppendString(config.Zap.Prefix + t.Format("2006/01/02 - 15:04:05"))
}
