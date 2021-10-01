package validation

import (
	"fmt"
	"reflect"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func Equals(t interface{}) validation.Rule {
	return By(equals(t))
}

// equals is a custom validator to validate current object's value to be deeply-equal to the given value.
func equals(t interface{}) validation.RuleFunc {
	return func(value interface{}) error {
		if t != value {
			return validation.NewError("", fmt.Sprintf("must equals '%+v' (%v)", t, reflect.TypeOf(t)))
		}

		return nil
	}
}
