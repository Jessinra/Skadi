package validation

import validation "github.com/go-ozzo/ozzo-validation/v4"

// mapFieldRules is a method grouping of rules specifying required/optional keys name and it's value type.
type mapFieldRules struct{}

var M = mapFieldRules{}

func (mapFieldRules) KeyRequired(key interface{}, rules ...validation.Rule) *validation.KeyRules {
	rules = append(rules, validation.Required)
	return validation.Key(key, rules...)
}

func (mapFieldRules) BoolRequired(key interface{}, rules ...validation.Rule) *validation.KeyRules {
	rules = append(rules, validation.Required, IsBool)
	return validation.Key(key, rules...)
}

func (mapFieldRules) FloatRequired(key interface{}, rules ...validation.Rule) *validation.KeyRules {
	rules = append(rules, validation.Required, IsFloat)
	return validation.Key(key, rules...)
}

func (mapFieldRules) StringRequired(key interface{}, rules ...validation.Rule) *validation.KeyRules {
	rules = append(rules, validation.Required, IsString)
	return validation.Key(key, rules...)
}

func (mapFieldRules) DateRequired(key interface{}, rules ...validation.Rule) *validation.KeyRules {
	rules = append(rules, validation.Required, IsDate)
	return validation.Key(key, rules...)
}

func (mapFieldRules) TimestampRequired(key interface{}, rules ...validation.Rule) *validation.KeyRules {
	rules = append(rules, validation.Required, IsRFC3339)
	return validation.Key(key, rules...)
}

func (mapFieldRules) Key(key interface{}, rules ...validation.Rule) *validation.KeyRules {
	return validation.Key(key, rules...).Optional()
}

func (mapFieldRules) Bool(key interface{}, rules ...validation.Rule) *validation.KeyRules {
	rules = append(rules, IsBool)
	return validation.Key(key, rules...).Optional()
}

func (mapFieldRules) Float(key interface{}, rules ...validation.Rule) *validation.KeyRules {
	rules = append(rules, IsFloat)
	return validation.Key(key, rules...).Optional()
}

func (mapFieldRules) String(key interface{}, rules ...validation.Rule) *validation.KeyRules {
	rules = append(rules, IsString)
	return validation.Key(key, rules...).Optional()
}

func (mapFieldRules) Date(key interface{}, rules ...validation.Rule) *validation.KeyRules {
	rules = append(rules, IsDate)
	return validation.Key(key, rules...).Optional()
}

func (mapFieldRules) Timestamp(key interface{}, rules ...validation.Rule) *validation.KeyRules {
	rules = append(rules, IsRFC3339)
	return validation.Key(key, rules...).Optional()
}
