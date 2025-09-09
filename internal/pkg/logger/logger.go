package logger

import (
	"context"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/coin50etf/coin-market/internal/pkg/config"
	"github.com/coin50etf/coin-market/internal/pkg/utils/ctxutils"
)

var logger *zap.SugaredLogger

func InitLogger() {
	logConfig := config.Conf.Log
	// 解析日志级别
	var logLevel zapcore.Level
	if err := logLevel.UnmarshalText([]byte(logConfig.Level)); err != nil {
		logLevel = zapcore.InfoLevel // 默认 Info 级别
	}

	// 配置日志格式
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		CallerKey:      "caller",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		EncodeTime:     zapcore.TimeEncoderOfLayout(time.RFC3339),
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
	}

	// 文件日志配置（lumberjack 负责日志切割）
	/*fileSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   logConfig.LogPath,
		MaxSize:    logConfig.MaxSize,
		MaxBackups: logConfig.MaxBackups,
		MaxAge:     logConfig.MaxAge,
		Compress:   true,
	})*/

	// 标准输出
	consoleSyncer := zapcore.AddSync(os.Stdout)

	// 构造 core
	core := zapcore.NewTee(
		zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConfig), consoleSyncer, logLevel),
		//zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), fileSyncer, logLevel),
	)

	logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1)).Sugar()
}

func Debug(ctx context.Context, msg string, fields ...interface{}) {
	withContext(ctx).Debugw(msg, fields...)
}

func Info(ctx context.Context, msg string, fields ...interface{}) {
	withContext(ctx).Infow(msg, fields...)
}

func Warn(ctx context.Context, msg string, fields ...interface{}) {
	withContext(ctx).Warnw(msg, fields...)
}

func Error(ctx context.Context, msg string, fields ...interface{}) {
	withContext(ctx).Errorw(msg, fields...)
}

func Fatal(ctx context.Context, msg string, fields ...interface{}) {
	withContext(ctx).Fatalw(msg, fields...)
}

func withContext(ctx context.Context) *zap.SugaredLogger {
	traceID := ctxutils.GetTraceID(ctx)
	if traceID == "" {
		traceID = "unknown trace id"
	}
	userID := ctxutils.GetUserID(ctx)
	if userID == "" {
		userID = "unknown user id"
	}

	return logger.With("trace_id", traceID, "user_id", userID)
}
