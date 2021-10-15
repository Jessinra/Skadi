package app

import (
	"encoding/json"
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gitlab.com/trivery-id/skadi/utils/errors"
	"gitlab.com/trivery-id/skadi/utils/jwt"
	"gitlab.com/trivery-id/skadi/utils/metadata"
	writer "gitlab.com/trivery-id/skadi/utils/response-writer"
)

var errInvalidCredentials = errors.NewUnauthorizedError("invalid credentials")

type middleware func(next gin.HandlerFunc) gin.HandlerFunc

func chainMiddleware(mw ...middleware) middleware {
	return func(final gin.HandlerFunc) gin.HandlerFunc {
		return func(c *gin.Context) {
			last := final
			for i := len(mw) - 1; i >= 0; i-- {
				last = mw[i](last)
			}
			last(c)
		}
	}
}

func corsMiddleware() gin.HandlerFunc {
	allowedHost := map[string][]string{
		"dev": {
			"http://localhost:5000",
			"https://dev.trivery.id",
		},
	}

	return cors.New(cors.Config{
		AllowOrigins: allowedHost[os.Getenv("ENV")],
		AllowMethods: []string{
			"OPTIONS", "GET", "POST", "PUT", "PATCH",
		},
		AllowHeaders: []string{
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
			"accept",
			"origin",
		},
		AllowCredentials: true,
		MaxAge:           1 * time.Hour,
	})
}

func addUUIDToRequestCtxMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := metadata.NewContextWithUUID(c.Request.Context())
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func authenticatedUser(next gin.HandlerFunc) gin.HandlerFunc {
	m := chainMiddleware(parseJWT)
	return m(next)
}

func parseJWT(next gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		jwtToken := getBearerAuthToken(c)
		if jwtToken == "" {
			writer.WriteFailResponseFromError(c, errInvalidCredentials)
			return
		}

		claims, err := jwt.ParseToken(jwtToken)
		if err != nil {
			writer.WriteFailResponseFromError(c, errInvalidCredentials)
			return
		}

		userMeta := metadata.User{}
		bytes, _ := json.Marshal(claims)
		_ = json.Unmarshal(bytes, &userMeta)

		ctx = metadata.NewContextFromUser(ctx, userMeta)
		c.Request = c.Request.WithContext(ctx)
		next(c)
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
