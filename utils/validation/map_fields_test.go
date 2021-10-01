package validation_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/trivery-id/skadi/utils/validation"
)

func TestMapFieldRules(t *testing.T) {
	type testStruct struct {
		M map[string]interface{}
	}

	t.Run("ok - valid map, all required and optional given", func(t *testing.T) {
		st := testStruct{
			M: map[string]interface{}{
				"req_key":       "req_value",
				"req_bool":      true,
				"req_float":     3.14,
				"req_string":    "string",
				"req_date":      "2016-01-01",
				"req_timestamp": "2016-01-01T00:00:00Z",
				"ops_key":       "ops_value",
				"ops_bool":      true,
				"ops_float":     3.14,
				"ops_string":    "string",
				"ops_date":      "2016-01-01",
				"ops_timestamp": "2016-01-01T00:00:00Z",
			},
		}

		err := validation.ValidateStruct(&st,
			validation.Field(&st.M, validation.DynamicMap(
				validation.M.KeyRequired("req_key"),
				validation.M.BoolRequired("req_bool"),
				validation.M.FloatRequired("req_float"),
				validation.M.StringRequired("req_string"),
				validation.M.DateRequired("req_date"),
				validation.M.TimestampRequired("req_timestamp"),
				validation.M.Key("ops_key"),
				validation.M.Bool("ops_bool"),
				validation.M.Float("ops_float"),
				validation.M.String("ops_string"),
				validation.M.Date("ops_date"),
				validation.M.Timestamp("ops_timestamp"),
			)),
		)

		assert.Nil(t, err)
	})

	t.Run("ok - valid map, all optionals key", func(t *testing.T) {
		st := testStruct{
			M: map[string]interface{}{},
		}

		err := validation.ValidateStruct(&st,
			validation.Field(&st.M, validation.DynamicMap(
				validation.M.Key("ops_key"),
				validation.M.Bool("ops_bool"),
				validation.M.Float("ops_float"),
				validation.M.String("ops_string"),
				validation.M.Date("ops_date"),
				validation.M.Timestamp("ops_timestamp"),
			)),
		)

		assert.Nil(t, err)
	})

	t.Run("err - missing required keys", func(t *testing.T) {
		st := testStruct{
			M: map[string]interface{}{},
		}

		err := validation.ValidateStruct(&st,
			validation.Field(&st.M, validation.DynamicMap(
				validation.M.KeyRequired("req_key"),
				validation.M.BoolRequired("req_bool"),
				validation.M.FloatRequired("req_float"),
				validation.M.StringRequired("req_string"),
				validation.M.DateRequired("req_date"),
				validation.M.TimestampRequired("req_timestamp"),
			)),
		)

		assert.NotNil(t, err)
		assert.EqualError(t, err, "M: (req_bool: required key is missing; req_date: required key is missing; req_float: required key is missing; req_key: required key is missing; req_string: required key is missing; req_timestamp: required key is missing.).")
	})

	t.Run("err - value given, invalid type", func(t *testing.T) {
		st := testStruct{
			M: map[string]interface{}{
				"req_bool":      "true",
				"req_float":     "3.14",
				"req_string":    123,
				"req_date":      123,
				"req_timestamp": 123,
				"ops_bool":      "true",
				"ops_float":     "3.14",
				"ops_string":    123,
				"ops_date":      123,
				"ops_timestamp": 123,
			},
		}

		err := validation.ValidateStruct(&st,
			validation.Field(&st.M, validation.DynamicMap(
				validation.M.BoolRequired("req_bool"),
				validation.M.FloatRequired("req_float"),
				validation.M.StringRequired("req_string"),
				validation.M.DateRequired("req_date"),
				validation.M.TimestampRequired("req_timestamp"),
				validation.M.Bool("ops_bool"),
				validation.M.Float("ops_float"),
				validation.M.String("ops_string"),
				validation.M.Date("ops_date"),
				validation.M.Timestamp("ops_timestamp"),
			)),
		)

		assert.NotNil(t, err)
		assert.EqualError(t, err, "M: (ops_bool: must be boolean; ops_date: must follow ISO8601 format; ops_float: must be float32 or float64; ops_string: must be a string; ops_timestamp: must follow RFC3339 format; req_bool: must be boolean; req_date: must follow ISO8601 format; req_float: must be float32 or float64; req_string: must be a string; req_timestamp: must follow RFC3339 format.).")
	})
}
