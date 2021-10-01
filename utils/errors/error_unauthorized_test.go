package errors_test

import (
	"net/http"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	. "gitlab.com/trivery-id/skadi/utils/errors"
)

const errMsgUnauthorized = "it's an unauthorized error"

func Test_NewUnauthorizedError(t *testing.T) {
	t.Run("ok - created an unauthorized error", func(t *testing.T) {
		got := NewUnauthorizedError(errMsgUnauthorized)

		assert.NotNil(t, got)
		assert.True(t, IsUnauthorizedError(got))
		assert.Equal(t, http.StatusUnauthorized, got.Code())
		assert.EqualError(t, got, errMsgUnauthorized)
	})

	t.Run("ok - created an unauthorized error with original error", func(t *testing.T) {
		err := errors.New("original error")
		got := NewUnauthorizedError(errMsgUnauthorized, err)

		assert.NotNil(t, got)
		assert.True(t, IsUnauthorizedError(got))
		assert.Equal(t, http.StatusUnauthorized, got.Code())
		assert.Equal(t, errMsgUnauthorized, got.Message())
		assert.EqualError(t, got, "original error")
	})
}

func Test_IsUnauthorizedError(t *testing.T) {
	t.Run("ok - return true on NewUnauthorizedError", func(t *testing.T) {
		got := NewUnauthorizedError(errMsgUnauthorized)
		assert.True(t, IsUnauthorizedError(got))
	})

	t.Run("ok - return true on custom error with code 401", func(t *testing.T) {
		got := NewCustomError(http.StatusUnauthorized, errMsgUnauthorized)
		assert.True(t, IsUnauthorizedError(got))
	})
}
