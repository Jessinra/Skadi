package errors_test

import (
	"net/http"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	. "gitlab.com/trivery-id/skadi/utils/errors"
)

const errMsgForbidden = "it's a forbidden error"

func Test_NewForbiddenError(t *testing.T) {
	t.Run("ok - created a forbidden error", func(t *testing.T) {
		got := NewForbiddenError(errMsgForbidden)

		assert.NotNil(t, got)
		assert.True(t, IsForbiddenError(got))
		assert.Equal(t, http.StatusForbidden, got.Code())
		assert.EqualError(t, got, errMsgForbidden)
	})

	t.Run("ok - created a forbidden error with original error", func(t *testing.T) {
		err := errors.New("original error")
		got := NewForbiddenError(errMsgForbidden, err)

		assert.NotNil(t, got)
		assert.True(t, IsForbiddenError(got))
		assert.Equal(t, http.StatusForbidden, got.Code())
		assert.Equal(t, errMsgForbidden, got.Message())
		assert.EqualError(t, got, "original error")
	})
}

func Test_IsForbiddenError(t *testing.T) {
	t.Run("ok - return true on NewForbiddenError", func(t *testing.T) {
		got := NewForbiddenError(errMsgForbidden)
		assert.True(t, IsForbiddenError(got))
	})

	t.Run("ok - return true on custom error with code 403", func(t *testing.T) {
		got := NewCustomError(http.StatusForbidden, errMsgForbidden)
		assert.True(t, IsForbiddenError(got))
	})
}
