package aws_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/trivery-id/skadi/external/secret-manager"
	. "gitlab.com/trivery-id/skadi/external/secret-manager/aws"
)

const (
	testSecretID = "parnassus-test" // nolint gosec potential hardcoded credentials
	testRegion   = "ap-southeast-1"
)

func Test_awsSecretManager_LoadSecret_Integration(t *testing.T) {
	type testSecret struct {
		String string  `json:"k_string"`
		Int    int     `json:"k_int"`
		Float  float64 `json:"k_float"`
	}

	manager := NewSecretManager()

	t.Run("ok - loaded all secret", func(t *testing.T) {
		secretValues := &testSecret{}
		err := manager.LoadSecret(testSecretID, secretValues,
			secret.WithLocation(testRegion),
		)

		assert.Nil(t, err)
		assert.Equal(t, "value_01", secretValues.String)
		assert.Equal(t, 2, secretValues.Int)
		assert.Equal(t, 3.14, secretValues.Float)
	})

	t.Run("err - secret doesnt exists", func(t *testing.T) {
		secretValues := &testSecret{}
		err := manager.LoadSecret("testSecretID", secretValues,
			secret.WithLocation(testRegion),
		)

		assert.NotNil(t, err)
	})

	t.Run("err - secret doesnt exists on region", func(t *testing.T) {
		secretValues := &testSecret{}
		err := manager.LoadSecret(testSecretID, secretValues,
			secret.WithLocation("alaska"),
		)

		assert.NotNil(t, err)
	})
}
