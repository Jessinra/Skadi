package logger

import (
	"context"

	"go.uber.org/zap"
)

// With creates a child logger and adds structured fields to it.
// Fields added to the child don't affect the parent, and vice versa.
func (l *Logger) With(fields ...zap.Field) *Logger {
	return &Logger{
		zapLogger: l.zapLogger.With(fields...),
	}
}

// WithCtx creates a child logger and adds structured context to it.
// Fields added to the child don't affect the parent, and vice versa.
// Context values is parsed using the default context parser provided.
func (l *Logger) WithCtx(ctx context.Context) *Logger {
	if defaultContextParser == nil {
		l.Warn("logger: unable to parse context: default context parser is not configured properly, use logger.SetDefaultContextParser() to set the default parser")
		return l
	}

	ctxFields := defaultContextParser.Parse(ctx)

	zapFields := []zap.Field{}
	for k, v := range ctxFields {
		zapFields = append(zapFields, zap.Any(k, v))
	}

	return &Logger{
		zapLogger: l.zapLogger.With(zapFields...),
	}
}

func (l *Logger) Debug(msg string, tags ...zap.Field) {
	l.zapLogger.Debug(msg, tags...)
	_ = l.zapLogger.Sync()
}

func (l *Logger) Info(msg string, tags ...zap.Field) {
	l.zapLogger.Info(msg, tags...)
	_ = l.zapLogger.Sync()
}

func (l *Logger) Warn(msg string, tags ...zap.Field) {
	l.zapLogger.Warn(msg, tags...)
	_ = l.zapLogger.Sync()
}

func (l *Logger) Error(msg string, tags ...zap.Field) {
	l.zapLogger.Error(msg, tags...)
	_ = l.zapLogger.Sync()
}
