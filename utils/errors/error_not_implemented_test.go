package errors_test

import (
	"net/http"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	. "gitlab.com/trivery-id/skadi/utils/errors"
)

const errMsgNotImplemented = "it's a not implemented error"

func Test_NewNotImplementedError(t *testing.T) {
	t.Run("ok - created a not implemented error", func(t *testing.T) {
		got := NewNotImplementedError(errMsgNotImplemented)

		assert.NotNil(t, got)
		assert.True(t, IsNotImplementedError(got))
		assert.Equal(t, http.StatusNotImplemented, got.Code())
		assert.EqualError(t, got, errMsgNotImplemented)
	})

	t.Run("ok - created a not implemented error with original error", func(t *testing.T) {
		err := errors.New("original error")
		got := NewNotImplementedError(errMsgNotImplemented, err)

		assert.NotNil(t, got)
		assert.True(t, IsNotImplementedError(got))
		assert.Equal(t, http.StatusNotImplemented, got.Code())
		assert.Equal(t, errMsgNotImplemented, got.Message())
		assert.EqualError(t, got, "original error")
	})
}

func Test_IsNotImplementedError(t *testing.T) {
	t.Run("ok - return true on NewNotImplementedError", func(t *testing.T) {
		got := NewNotImplementedError(errMsgNotImplemented)
		assert.True(t, IsNotImplementedError(got))
	})

	t.Run("ok - return true on custom error with code 501", func(t *testing.T) {
		got := NewCustomError(http.StatusNotImplemented, errMsgNotImplemented)
		assert.True(t, IsNotImplementedError(got))
	})
}
