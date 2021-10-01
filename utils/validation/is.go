package validation

import (
	"fmt"
	"math"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

var (
	IsAlpha       = is.Alpha
	IsE164        = is.E164
	IsEmailFormat = is.EmailFormat
	IsIPv4        = is.IPv4
	IsUUIDv4      = is.UUIDv4
)

var (
	IsBool           = By(isBool)
	IsString         = By(isString)
	IsFloat          = By(isFloat)
	IsFloatLatitude  = By(isFloatLatitude)
	IsFloatLongitude = By(isFloatLongitude)
	IsInt            = By(isInt)
	IsNumber         = By(isNumber)
	IsRoundNumber    = By(isRoundNumber)
	IsRFC3339        = By(isRFC3339)
	IsDate           = By(isISO8601Date)
	IsPositiveOrZero = By(isPositiveOrZero)
	IsStrongPassword = By(isStrongPassword)
)

func isBool(value interface{}) error {
	switch value.(type) {
	case bool:
	default:
		return validation.NewError("", "must be boolean")
	}

	return nil
}

func isString(value interface{}) error {
	switch value.(type) {
	case string:
	default:
		return validation.NewError("", "must be a string")
	}

	return nil
}

func isFloat(value interface{}) error {
	switch value.(type) {
	case float32:
	case float64:
	default:
		return validation.NewError("", "must be float32 or float64")
	}

	return nil
}

func isFloatLatitude(value interface{}) error {
	if err := isFloat(value); err != nil {
		return err
	}

	valueStr := fmt.Sprintf("%v", value)
	return validation.Validate(valueStr, validation.Required, is.Latitude)
}

func isFloatLongitude(value interface{}) error {
	if err := isFloat(value); err != nil {
		return err
	}

	valueStr := fmt.Sprintf("%v", value)
	return validation.Validate(valueStr, validation.Required, is.Longitude)
}

func isInt(value interface{}) error {
	switch value.(type) {
	case int:
	case int8:
	case int16:
	case int32:
	case int64:
	default:
		return validation.NewError("", "must be an integer")
	}

	return nil
}

func isNumber(value interface{}) error {
	if err := isFloat(value); err == nil {
		return nil
	}
	if err := isInt(value); err == nil {
		return nil
	}

	return validation.NewError("", "must be a number")
}

// isRoundNumber validate if the value is uint, int, or float without decimal points.
func isRoundNumber(value interface{}) error {
	switch v := value.(type) {
	case uint, uint8, uint16, uint32, uint64:
		return nil
	case int, int8, int16, int32, int64:
		return nil
	case float32:
		if _, f := math.Modf(float64(v)); f == 0 {
			return nil
		}
	case float64:
		if _, f := math.Modf(v); f == 0 {
			return nil
		}
	}

	return validation.NewError("", "must be an uint, int, or float without decimal points")
}

// isRFC3339 is a custom validator to validate if the value follow timestamp RFC3339 format.
func isRFC3339(value interface{}) error {
	val, _ := value.(string)
	if _, err := time.Parse(time.RFC3339, val); err != nil {
		return validation.NewError("", "must follow RFC3339 format")
	}

	return nil
}

// isISO8601Date is a custom validator to validate if the value follow ISO8601 date format (YYYY-MM-DD).
func isISO8601Date(value interface{}) error {
	val, _ := value.(string)
	if _, err := time.Parse("2006-01-02", val); err != nil {
		return validation.NewError("", "must follow ISO8601 format")
	}

	return nil
}

func isPositiveOrZero(value interface{}) error {
	isPositiveOrZero := false
	switch v := value.(type) {
	case uint, uint8, uint16, uint32, uint64:
		isPositiveOrZero = true
	case int:
		isPositiveOrZero = v >= 0
	case int8:
		isPositiveOrZero = v >= 0
	case int16:
		isPositiveOrZero = v >= 0
	case int32:
		isPositiveOrZero = v >= 0
	case int64:
		isPositiveOrZero = v >= 0
	case float32:
		isPositiveOrZero = v >= 0
	case float64:
		isPositiveOrZero = v >= 0
	default:
		return validation.NewError("", "must be a number")
	}

	if !isPositiveOrZero {
		return validation.NewError("", "must be positive or zero value")
	}

	return nil
}

func isStrongPassword(value interface{}) error {
	// TODO: impelement later
	if err := validation.Validate(value, validation.Length(8, 64)); err != nil { // nolint
		return validation.NewError("", "must be at least 8 chars and not longer than 64 chars")
	}

	return nil
}
