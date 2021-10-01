package errors_test

import (
	"net/http"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	. "gitlab.com/trivery-id/skadi/utils/errors"
)

const errMsgNotFound = "it's a not found"

func Test_NewNotFoundError(t *testing.T) {
	t.Run("ok - created a not found error", func(t *testing.T) {
		got := NewNotFoundError(errMsgNotFound)

		assert.NotNil(t, got)
		assert.True(t, IsNotFoundError(got))
		assert.Equal(t, http.StatusNotFound, got.Code())
		assert.EqualError(t, got, errMsgNotFound)
	})

	t.Run("ok - created a not found error with original error", func(t *testing.T) {
		err := errors.New("original error")
		got := NewNotFoundError(errMsgNotFound, err)

		assert.NotNil(t, got)
		assert.True(t, IsNotFoundError(got))
		assert.Equal(t, http.StatusNotFound, got.Code())
		assert.Equal(t, errMsgNotFound, got.Message())
		assert.EqualError(t, got, "original error")
	})
}

func Test_IsNotFoundError(t *testing.T) {
	t.Run("ok - return true on NewNotFoundError", func(t *testing.T) {
		got := NewNotFoundError(errMsgNotFound)
		assert.True(t, IsNotFoundError(got))
	})

	t.Run("ok - return true on custom error with code 404", func(t *testing.T) {
		got := NewCustomError(http.StatusNotFound, errMsgNotFound)
		assert.True(t, IsNotFoundError(got))
	})
}
