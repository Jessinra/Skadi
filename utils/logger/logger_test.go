package logger_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/trivery-id/skadi/utils/logger"
	"gitlab.com/trivery-id/skadi/utils/metadata"
	"go.uber.org/zap"
)

var errAny = errors.New("hehe")

func TestExampleLogging(t *testing.T) {
	t.Run("log using default logger", func(t *testing.T) {
		assert.NotPanics(t, func() {
			logger.Info("test log info")
			logger.Warn("test log warning")
			logger.Error("test log error", errAny)
		})
	})
}

func TestLoggerWithAdditionalContext(t *testing.T) {
	t.Run("log using logger with additional context", func(t *testing.T) {
		log := logger.GetLogger().With(
			zap.String("string", "STRING"),
			zap.Int("int", 123),
			zap.Float64("float", 3.14),
		)

		assert.NotPanics(t, func() {
			log.Info("test log info with context")
			log.Warn("test log warn with context")
			log.Error("test log error with context")
		})
	})
}

func TestLoggerWithCtx(t *testing.T) {
	// Notes: run test on debug mode to see the log

	logger.SetDefaultContextParser(metadata.LoggerContextparser{})

	t.Run("log using logger with ctx (user metadata) and added other additional context", func(t *testing.T) {
		userMetadata := metadata.UserMetadata{
			ID:       9342,
			ClientID: 2131,
		}

		ctxWithUser := metadata.NewContextFromUser(context.Background(), userMetadata)
		log := logger.GetLoggerWithCtx(ctxWithUser).With(
			zap.String("string", "STRING"),
			zap.Int("int", 123),
			zap.Float64("float", 3.14),
		)

		assert.NotPanics(t, func() {
			log.Info("test log info with context")
			log.Warn("test log warn with context")
			log.Error("test log error with context")
		})
	})

	t.Run("log using logger with ctx (api key metadata) and added other additional context", func(t *testing.T) {
		apiKeyMetadata := metadata.APIKeyMetadata{
			Key: "ablksjdlfkajsldfkaablksjdlfkajsldfkajslkjjslkj",
		}

		ctxWithAPIKey := metadata.NewContextFromAPIKey(context.Background(), apiKeyMetadata)
		log := logger.GetLoggerWithCtx(ctxWithAPIKey).With(
			zap.String("string", "STRING"),
			zap.Int("int", 123),
			zap.Float64("float", 3.14),
		)

		assert.NotPanics(t, func() {
			log.Info("test log info with context")
			log.Warn("test log warn with context")
			log.Error("test log error with context")
		})
	})
}
