package validation

import validation "github.com/go-ozzo/ozzo-validation/v4"

// This is the 'extension' package of "github.com/go-ozzo/ozzo-validation/v4",
// consist of helpers to simplify the schema definition process.

var (
	Required = validation.Required
	NotNil   = validation.NotNil
)

func ValidateStruct(structPtr interface{}, fields ...*validation.FieldRules) error {
	return validation.ValidateStruct(structPtr, fields...)
}

func Validate(value interface{}, rules ...validation.Rule) error {
	return validation.Validate(value, rules...)
}

func Field(fieldPtr interface{}, rules ...validation.Rule) *validation.FieldRules {
	return validation.Field(fieldPtr, rules...)
}

func By(f validation.RuleFunc) validation.Rule {
	return validation.By(f)
}

func In(values ...interface{}) validation.InRule {
	return validation.In(values...)
}

func Each(rules ...validation.Rule) validation.EachRule {
	return validation.Each(rules...)
}

func NewError(code, message string) validation.Error {
	return validation.NewError(code, message)
}

func Min(min interface{}) validation.ThresholdRule {
	return validation.Min(min)
}

func Max(max interface{}) validation.ThresholdRule {
	return validation.Max(max)
}

// GreaterOrEqual alias to validation.Min(...) to enable better readability.
func GreaterOrEqual(value interface{}) validation.ThresholdRule {
	return validation.Min(value)
}

// LessOrEqual alias to validation.Max(...) to enable better readability.
func LessOrEqual(value interface{}) validation.ThresholdRule {
	return validation.Max(value)
}

func Length(min, max int) validation.LengthRule {
	return validation.Length(min, max)
}
