package errors_test

import (
	stderrors "errors"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	. "gitlab.com/trivery-id/skadi/utils/errors"
)

//nolint //errorlint
func Test_New(t *testing.T) {
	t.Run("ok - created standard error", func(t *testing.T) {
		errMsg := "random err message"
		err := stderrors.New(errMsg)

		got := New(errMsg)

		assert.Equal(t, err, got)
		assert.EqualError(t, got, errMsg)

		_, isStandardError := got.(error)
		assert.True(t, isStandardError, "should be a standard error")
	})
}

func Test_Is(t *testing.T) {
	originalErr := errors.New("someError")

	t.Run("ok - return true on exact same error", func(t *testing.T) {
		got := Is(originalErr, originalErr)
		assert.True(t, got)
	})

	t.Run("ok - return true on wrapped original error", func(t *testing.T) {
		wrappedErr := errors.Wrap(originalErr, "some wrapper")

		got := Is(wrappedErr, originalErr)
		assert.True(t, got)
	})

	t.Run("ok - return false on different error", func(t *testing.T) {
		anotherErr := errors.New("not original")

		got := Is(anotherErr, originalErr)
		assert.False(t, got)
	})
}

func Test_Wrap(t *testing.T) {
	originalErr := errors.New("someError")
	got := Wrap(originalErr, "some wrapper")

	t.Run("ok - added wrapper to error message", func(t *testing.T) {
		assert.EqualError(t, got, "some wrapper: someError")
	})

	t.Run("ok - doesn't change internal error type / value", func(t *testing.T) {
		assert.True(t, errors.Is(got, originalErr))
	})

	t.Run("ok - calling cause return the original error", func(t *testing.T) {
		assert.Equal(t, originalErr, errors.Cause(got))
	})
}

func Test_Cause(t *testing.T) {
	originalErr := errors.New("someError")
	wrappedErr := errors.Wrap(originalErr, "some wrapper")

	t.Run("ok - calling cause return the original error", func(t *testing.T) {
		got := Cause(wrappedErr)
		assert.Equal(t, originalErr, got)
	})

	t.Run("ok - calling cause on non wrapped error return the error (doesn't do anything)", func(t *testing.T) {
		got := Cause(originalErr)
		assert.Equal(t, originalErr, got)
	})
}
