package uuid

import (
	"github.com/google/uuid"
	"gitlab.com/trivery-id/skadi/utils/logger"
	"go.uber.org/zap"
)

func NewUUID() string {
	defer func() {
		if r := recover(); r != nil {
			logger.Warn("failed to generate UUID",
				zap.String("function", "uuid.NewUUID"),
			)
		}
	}()

	return uuid.New().String()
}
