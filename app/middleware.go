package app

import (
	"strings"

	"github.com/gin-gonic/gin"
	"gitlab.com/trivery-id/skadi/utils/errors"
	writer "gitlab.com/trivery-id/skadi/utils/response-writer"
)

var errInvalidClientCredentials = errors.NewUnauthorizedError("invalid credentials")

func parseJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		jwtToken := getBearerAuthToken(c)
		if jwtToken == "" {
			writer.WriteFailResponseFromError(c, errInvalidClientCredentials)
			return
		}

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

// getBearerAuthToken parse Bearer token from Authorization header.
func getBearerAuthToken(c *gin.Context) (bearerToken string) {
	const prefix = "Bearer"

	auth := c.Request.Header.Get("Authorization")
	if auth != "" && strings.HasPrefix(auth, prefix) {
		bearerToken = auth[len(prefix):]
	}

	return strings.TrimSpace(bearerToken)
}
