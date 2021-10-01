package logger

import (
	"context"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var defaultLogger *Logger

type Logger struct {
	zapLogger *zap.Logger
}

// init ensure all test code with logger as it's dependency wont panic.
func init() {
	InitLogger()
}

func InitLogger() {
	zapLogConfig := zap.Config{
		OutputPaths: []string{"stdout"},
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "msg",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	zapLogger, err := zapLogConfig.Build()
	if err != nil {
		panic(err)
	}

	defaultLogger = &Logger{
		zapLogger: zapLogger,
	}
}

func GetLogger() *Logger {
	return defaultLogger
}

func GetLoggerWithCtx(ctx context.Context) *Logger {
	return defaultLogger.WithCtx(ctx)
}

func Debug(msg string, tags ...zap.Field) {
	defaultLogger.Debug(msg, tags...)
}

func Info(msg string, tags ...zap.Field) {
	defaultLogger.Info(msg, tags...)
}

func Warn(msg string, tags ...zap.Field) {
	defaultLogger.Warn(msg, tags...)
}

func Error(msg string, err error, tags ...zap.Field) {
	if err != nil {
		tags = append(tags, zap.NamedError("error", err))
	}

	defaultLogger.Error(msg, tags...)
}

func DebugWithCtx(ctx context.Context, msg string, tags ...zap.Field) {
	defaultLogger.WithCtx(ctx).Debug(msg, tags...)
}

func InfoWithCtx(ctx context.Context, msg string, tags ...zap.Field) {
	defaultLogger.WithCtx(ctx).Info(msg, tags...)
}

func WarnWithCtx(ctx context.Context, msg string, tags ...zap.Field) {
	defaultLogger.WithCtx(ctx).Warn(msg, tags...)
}

func ErrorWithCtx(ctx context.Context, msg string, err error, tags ...zap.Field) { // nolint argument-limit
	if err != nil {
		tags = append(tags, zap.NamedError("error", err))
	}

	defaultLogger.WithCtx(ctx).Error(msg, tags...)
}
