package writer

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.com/trivery-id/skadi/utils/logger"
	"gitlab.com/trivery-id/skadi/utils/metadata"
	"go.uber.org/zap"
)

func WriteSuccessResponse(c *gin.Context, statusCode int, data interface{}) {
	c.AbortWithStatusJSON(statusCode, data)
}

func WriteFailResponse(c *gin.Context, resp ErrorResponse) {
	c.AbortWithStatusJSON(resp.Status, resp)
}

func WriteFailResponseFromError(c *gin.Context, err error, opts ...Option) {
	errResp := NewErrorResponse(err)
	errResp.UUID = metadata.GetUUIDFromContext(c.Request.Context())
	errResp.Timestamp = time.Now()

	options := parseOptions(opts...)

	// Log all error on warning level
	if options.EnableErrorLogging {
		traceLogMsg := fmt.Sprintf("%s %s %s: %d %s", errResp.UUID, c.Request.Method, c.Request.URL.Path, errResp.Status, errResp.Message)
		logger.WarnWithCtx(c.Request.Context(), traceLogMsg, zap.Error(err))
	}

	c.AbortWithStatusJSON(errResp.Status, errResp)
}
