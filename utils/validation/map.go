package validation

import (
	"fmt"
	"reflect"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func Map(keys ...*validation.KeyRules) validation.MapRule {
	return validation.Map(keys...)
}

// DynamicMap decorates `validation.Map` to allow extra keys by default.
func DynamicMap(keys ...*validation.KeyRules) validation.MapRule {
	return validation.Map(keys...).AllowExtraKeys()
}

// RequireAtLeastOneKey is a custom validator to check whether current map have at least one of the specified keys.
// should return error if it's not a map, should skip checks if there're no keys given.
func RequireAtLeastOneKey(keys ...string) validation.Rule {
	return By(func(value interface{}) error {
		if len(keys) == 0 {
			return nil
		}

		v := reflect.ValueOf(value)
		if v.Kind() != reflect.Map {
			return validation.NewError("", "must be a map")
		}

		m := map[string]bool{}
		for _, key := range v.MapKeys() {
			m[key.String()] = true
		}

		anyKeyExist := false
		for _, key := range keys {
			anyKeyExist = anyKeyExist || m[key]
		}

		if !anyKeyExist {
			errMsg := fmt.Sprintf("must have at least one of ['%s']", strings.Join(keys, "', '"))
			return validation.NewError("", errMsg)
		}

		return nil
	})
}

// Key validate a map's key existence and it's value using the given rules.
func Key(key interface{}, rules ...validation.Rule) *validation.KeyRules {
	return validation.Key(key, rules...)
}
