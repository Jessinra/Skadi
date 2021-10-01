package errors_test

import (
	"net/http"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	. "gitlab.com/trivery-id/skadi/utils/errors"
)

const (
	errCodeCustomError = http.StatusBadRequest
	errMsgCustomError  = "it's a custom (olympuspublic) error"
)

func Test_NewCustomError(t *testing.T) {
	t.Run("ok - created a custom (olympuspublic) error", func(t *testing.T) {
		got := NewCustomError(errCodeCustomError, errMsgCustomError)

		assert.NotNil(t, got)
		assert.True(t, IsCustomError(got))
		assert.Equal(t, errCodeCustomError, got.Code())
		assert.EqualError(t, got, errMsgCustomError)
	})

	t.Run("ok - created a custom error with original error", func(t *testing.T) {
		err := errors.New("original error")
		got := NewCustomError(errCodeCustomError, errMsgCustomError, err)

		assert.NotNil(t, got)
		assert.True(t, IsBadRequestError(got))
		assert.Equal(t, errCodeCustomError, got.Code())
		assert.Equal(t, errMsgCustomError, got.Message())
		assert.EqualError(t, got, "original error")
	})
}
