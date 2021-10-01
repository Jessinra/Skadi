package dockertest_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	. "gitlab.com/trivery-id/skadi/utils/dockertest"
)

func TestNewPostgreSQLPool_Integration(t *testing.T) {
	t.Run("ok - successfully created docker pool with postgresql", func(t *testing.T) {
		got, err := NewPostgreSQLPool()
		defer func() {
			_ = got.Purge()
		}()

		assert.NotNil(t, got)
		assert.NotEmpty(t, got.Credential.Host)
		assert.NotEmpty(t, got.Credential.Port)
		assert.Equal(t, "postgres", got.Credential.Username)
		assert.Equal(t, "postgres", got.Credential.Password)
		assert.Equal(t, "database", got.Credential.DBName)
		assert.Nil(t, err)
	})
}
