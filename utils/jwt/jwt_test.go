package jwt_test

import (
	"os"
	"testing"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	. "gitlab.com/trivery-id/skadi/utils/jwt"
)

func TestMain(m *testing.M) {
	SetSignKey("testJWTHelloWorld")
	os.Exit(m.Run())
}

func TestNewToken(t *testing.T) {
	t.Run("ok - successfully create new token", func(t *testing.T) {
		claims := jwtgo.MapClaims(map[string]interface{}{
			"hello": "world",
		})

		got, err := NewToken(claims)
		assert.Equal(t, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJoZWxsbyI6IndvcmxkIn0.b8XABTO5-J-5Rx-CHD0BsE71EYzNyzmEuufz0FSYpC0", got)
		assert.Nil(t, err)
	})
}

func TestParseToken(t *testing.T) {
	t.Run("ok - successfully parse and extract token", func(t *testing.T) {
		claims := map[string]interface{}{
			"hello": "world",
		}

		got, err := ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJoZWxsbyI6IndvcmxkIn0.b8XABTO5-J-5Rx-CHD0BsE71EYzNyzmEuufz0FSYpC0")
		assert.Equal(t, claims, got)
		assert.Nil(t, err)
	})

	t.Run("err - expired token", func(t *testing.T) {
		claims := jwtgo.MapClaims(map[string]interface{}{
			"exp": 5000,
		})

		expiredToken, err := NewToken(claims)
		if err != nil {
			t.Fatal(err.Error())
		}

		got, err := ParseToken(expiredToken)
		assert.Nil(t, got)
		assert.NotNil(t, err)
		assert.EqualError(t, err, "invalid JWT token")
	})
}
