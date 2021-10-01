package validation_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/trivery-id/skadi/utils/validation"
)

func TestRequireAtLeastOneKey(t *testing.T) {
	type testMap struct {
		InterfaceMap map[string]interface{}
		StringMap    map[string]string
		FloatMap     map[string]float64
	}

	t.Run("ok - has all of the required keys", func(t *testing.T) {
		requiredKeys := []string{"key_a", "key_b", "key_c"}
		tmap := testMap{
			InterfaceMap: map[string]interface{}{
				"key_a": "a",
				"key_b": "b",
				"key_c": "c",
			},
			StringMap: map[string]string{
				"key_a": "a",
				"key_b": "b",
				"key_c": "c",
			},
			FloatMap: map[string]float64{
				"key_a": 1.2,
				"key_b": 1.2,
				"key_c": 1.2,
			},
		}

		err := validation.ValidateStruct(&tmap,
			validation.Field(&tmap.InterfaceMap, validation.RequireAtLeastOneKey(requiredKeys...)),
			validation.Field(&tmap.StringMap, validation.RequireAtLeastOneKey(requiredKeys...)),
			validation.Field(&tmap.FloatMap, validation.RequireAtLeastOneKey(requiredKeys...)),
		)

		assert.Nil(t, err)
	})

	t.Run("ok - has one of the required keys", func(t *testing.T) {
		requiredKeys := []string{"key_a", "key_b", "key_c"}
		tmap := testMap{
			InterfaceMap: map[string]interface{}{
				"key_a": "a",
			},
			StringMap: map[string]string{
				"key_a": "a",
			},
			FloatMap: map[string]float64{
				"key_a": 1.2,
			},
		}

		err := validation.ValidateStruct(&tmap,
			validation.Field(&tmap.InterfaceMap, validation.RequireAtLeastOneKey(requiredKeys...)),
			validation.Field(&tmap.StringMap, validation.RequireAtLeastOneKey(requiredKeys...)),
			validation.Field(&tmap.FloatMap, validation.RequireAtLeastOneKey(requiredKeys...)),
		)

		assert.Nil(t, err)
	})

	t.Run("ok - no required keys (should ignore checks)", func(t *testing.T) {
		tmap := testMap{
			InterfaceMap: map[string]interface{}{
				"key_a": "a",
			},
			StringMap: map[string]string{
				"key_a": "a",
			},
			FloatMap: map[string]float64{
				"key_a": 1.2,
			},
		}

		err := validation.ValidateStruct(&tmap,
			validation.Field(&tmap.InterfaceMap, validation.RequireAtLeastOneKey()),
			validation.Field(&tmap.StringMap, validation.RequireAtLeastOneKey()),
			validation.Field(&tmap.FloatMap, validation.RequireAtLeastOneKey()),
		)

		assert.Nil(t, err)
	})

	t.Run("err - doesn't have any of the required keys", func(t *testing.T) {
		requiredKeys := []string{"key_d", "key_e", "key_f"}
		tmap := testMap{
			InterfaceMap: map[string]interface{}{
				"key_a": "a",
			},
			StringMap: map[string]string{
				"key_a": "a",
			},
			FloatMap: map[string]float64{
				"key_a": 1.2,
			},
		}

		err := validation.ValidateStruct(&tmap,
			validation.Field(&tmap.InterfaceMap, validation.RequireAtLeastOneKey(requiredKeys...)),
			validation.Field(&tmap.StringMap, validation.RequireAtLeastOneKey(requiredKeys...)),
			validation.Field(&tmap.FloatMap, validation.RequireAtLeastOneKey(requiredKeys...)),
		)

		assert.NotNil(t, err)
		assert.EqualError(t, err, "FloatMap: must have at least one of ['key_d', 'key_e', 'key_f']; InterfaceMap: must have at least one of ['key_d', 'key_e', 'key_f']; StringMap: must have at least one of ['key_d', 'key_e', 'key_f'].")
	})

	t.Run("err - value not a map", func(t *testing.T) {
		type testStructs struct {
			Interface interface{}
		}

		requiredKeys := []string{"key_a"}
		tStruct := testStructs{
			Interface: "string",
		}

		err := validation.ValidateStruct(&tStruct,
			validation.Field(&tStruct.Interface, validation.RequireAtLeastOneKey(requiredKeys...)),
		)

		assert.NotNil(t, err)
		assert.EqualError(t, err, "Interface: must be a map.")
	})
}
