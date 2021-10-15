package jwt_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	. "gitlab.com/trivery-id/skadi/utils/jwt"
)

func TestNewUnmaskTokenClaims(t *testing.T) {
	t.Run("ok - new created token is valid", func(t *testing.T) {
		claims := NewStandardClaims(WithExpiresAt(time.Now().Add(1 * time.Hour)))

		err := claims.Valid()
		assert.Nil(t, err)
	})

	t.Run("ok - default expire is 30 mins", func(t *testing.T) {
		claims := NewStandardClaims()

		defaultExpires := time.Now().Add(30 * time.Minute).UTC().Unix()
		assert.InEpsilon(t, defaultExpires, claims.ExpiresAt, 1)
	})

	t.Run("ok - expired token is invalid", func(t *testing.T) {
		claims := NewStandardClaims(WithExpiresAt(time.Now().Add(-1 * time.Hour)))

		err := claims.Valid()
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "token is expired")
	})
}
