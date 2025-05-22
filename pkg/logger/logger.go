package logger

import (
	"go_ecommerce_backend/pkg/setting"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type LoggerZap struct {
	*zap.Logger
}

func NewLogger(config setting.LoggerSetting) *LoggerZap {
	logLevel := config.Log_level

	var level zapcore.Level
	switch logLevel {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warn": // fixed: "warm" => "warn"
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	case "fatal":
		level = zapcore.FatalLevel
	default:
		level = zapcore.InfoLevel
	}

	encoder := getEncoderLog()

	hook := &lumberjack.Logger{
		Filename:   config.File_log_name, // fixed: typo "strorages" => "storages" in actual value
		MaxSize:    config.Max_size,      // in megabytes
		MaxBackups: config.Max_backups,
		MaxAge:     config.Max_age, // in days
		Compress:   config.Compress,
	}

	writer := zapcore.NewMultiWriteSyncer(
		zapcore.AddSync(os.Stdout),
		zapcore.AddSync(hook),
	)

	core := zapcore.NewCore(encoder, writer, level)

	return &LoggerZap{
		zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel)),
	}
}

func getEncoderLog() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()

	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encodeConfig.TimeKey = "Time"
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder

	return zapcore.NewJSONEncoder(encodeConfig)
}
