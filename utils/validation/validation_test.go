package validation_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gitlab.com/trivery-id/skadi/utils/validation"
)

func Test_Each(t *testing.T) {
	type Struct struct {
		Arr    []string
		ArrMap []map[string]string
	}

	t.Run("ok - all value valid", func(t *testing.T) {
		x := Struct{Arr: []string{"a", "b"}}
		err := validation.ValidateStruct(&x,
			validation.Field(&x.Arr, validation.Each(validation.IsAlpha)),
		)

		assert.Nil(t, err)
	})

	t.Run("ok - skip check if array is empty", func(t *testing.T) {
		x := Struct{Arr: []string{}}
		err := validation.ValidateStruct(&x,
			validation.Field(&x.Arr, validation.Each(validation.IsAlpha)),
		)

		assert.Nil(t, err)
	})

	t.Run("err - return error if any element fail the validation", func(t *testing.T) {
		x := Struct{Arr: []string{"a", "1"}}
		err := validation.ValidateStruct(&x,
			validation.Field(&x.Arr, validation.Each(validation.IsAlpha)),
		)

		assert.NotNil(t, err)
	})

	t.Run("ok - validate list of map", func(t *testing.T) {
		x := Struct{ArrMap: []map[string]string{
			{"val": "a"},
			{"val": "b"},
		}}
		err := validation.ValidateStruct(&x,
			validation.Field(&x.ArrMap, validation.Each(
				validation.DynamicMap(
					validation.M.String("val", validation.IsAlpha),
				),
			)),
		)

		assert.Nil(t, err)
	})

	t.Run("err - return error if any element fail the validation", func(t *testing.T) {
		x := Struct{ArrMap: []map[string]string{
			{"val": "a"},
			{"val": "1"},
		}}
		err := validation.ValidateStruct(&x,
			validation.Field(&x.ArrMap, validation.Each(
				validation.DynamicMap(
					validation.M.String("val", validation.IsAlpha),
				),
			)),
		)

		assert.NotNil(t, err)
		assert.EqualError(t, err, "ArrMap: (1: (val: must contain English letters only.).).")
	})
}

func Test_In(t *testing.T) {
	type Struct struct {
		Val string
	}

	t.Run("ok - value valid", func(t *testing.T) {
		x := Struct{Val: "a"}
		err := validation.ValidateStruct(&x,
			validation.Field(&x.Val, validation.In("a", "b")),
		)

		assert.Nil(t, err)
	})

	t.Run("ok - skip check if value is empty", func(t *testing.T) {
		x := Struct{Val: ""}
		err := validation.ValidateStruct(&x,
			validation.Field(&x.Val, validation.In("a", "b")),
		)

		assert.Nil(t, err)
	})

	t.Run("err - return error if value not in given list", func(t *testing.T) {
		x := Struct{Val: "z"}
		err := validation.ValidateStruct(&x,
			validation.Field(&x.Val, validation.In("a", "b")),
		)

		assert.NotNil(t, err)
	})
}

func Test_Min(t *testing.T) {
	type Struct struct {
		Float float64
		Int   int64
		Uint  uint64
		Time  time.Time
	}

	t.Run("ok - value valid", func(t *testing.T) {
		x := Struct{
			Float: 512.22,
			Int:   512,
			Uint:  512,
			Time:  time.Unix(1512000000, 0),
		}
		err := validation.ValidateStruct(&x,
			validation.Field(&x.Float, validation.Min(500.0)),
			validation.Field(&x.Int, validation.Min(500)),
			validation.Field(&x.Uint, validation.Min(uint64(500))),
			validation.Field(&x.Time, validation.Min(time.Unix(1502000000, 0))),
		)

		assert.Nil(t, err)
	})

	t.Run("ok - value equal min (minimum inclusive)", func(t *testing.T) {
		x := Struct{
			Float: 512.22,
			Int:   512,
			Uint:  512,
			Time:  time.Unix(1512000000, 0),
		}
		err := validation.ValidateStruct(&x,
			validation.Field(&x.Float, validation.Min(512.22)),
			validation.Field(&x.Int, validation.Min(512)),
			validation.Field(&x.Uint, validation.Min(uint64(512))),
			validation.Field(&x.Time, validation.Min(time.Unix(1512000000, 0))),
		)

		assert.Nil(t, err)
	})

	t.Run("err - return error if value less than minimum", func(t *testing.T) {
		x := Struct{
			Float: 512.22,
			Int:   512,
			Uint:  512,
			Time:  time.Unix(1512000000, 0),
		}
		err := validation.ValidateStruct(&x,
			validation.Field(&x.Float, validation.Min(1500.0)),
			validation.Field(&x.Int, validation.Min(1500)),
			validation.Field(&x.Uint, validation.Min(uint64(1500))),
		)
		errTime := validation.ValidateStruct(&x,
			validation.Field(&x.Time, validation.Min(time.Unix(1602000000, 0))),
		)

		assert.NotNil(t, err)
		assert.EqualError(t, err, "Float: must be no less than 1500; Int: must be no less than 1500; Uint: must be no less than 1500.")
		assert.Contains(t, errTime.Error(), "Time: must be no less than") // it shows local timezone on the error message
	})
}

func Test_Max(t *testing.T) {
	type Struct struct {
		Float float64
		Int   int64
		Uint  uint64
		Time  time.Time
	}

	t.Run("ok - value valid", func(t *testing.T) {
		x := Struct{
			Float: 512.22,
			Int:   512,
			Uint:  512,
			Time:  time.Unix(1512000000, 0),
		}
		err := validation.ValidateStruct(&x,
			validation.Field(&x.Float, validation.Max(1500.0)),
			validation.Field(&x.Int, validation.Max(1500)),
			validation.Field(&x.Uint, validation.Max(uint64(1500))),
			validation.Field(&x.Time, validation.Max(time.Unix(1602000000, 0))),
		)

		assert.Nil(t, err)
	})

	t.Run("ok - value equal max (maximum inclusive)", func(t *testing.T) {
		x := Struct{
			Float: 512.22,
			Int:   512,
			Uint:  512,
			Time:  time.Unix(1512000000, 0),
		}
		err := validation.ValidateStruct(&x,
			validation.Field(&x.Float, validation.Max(512.22)),
			validation.Field(&x.Int, validation.Max(512)),
			validation.Field(&x.Uint, validation.Max(uint64(512))),
			validation.Field(&x.Time, validation.Max(time.Unix(1512000000, 0))),
		)

		assert.Nil(t, err)
	})

	t.Run("err - return error if value is greater than maximum", func(t *testing.T) {
		x := Struct{
			Float: 512.22,
			Int:   512,
			Uint:  512,
			Time:  time.Unix(1512000000, 0),
		}
		err := validation.ValidateStruct(&x,
			validation.Field(&x.Float, validation.Max(500.0)),
			validation.Field(&x.Int, validation.Max(500)),
			validation.Field(&x.Uint, validation.Max(uint64(500))),
		)
		errTime := validation.ValidateStruct(&x,
			validation.Field(&x.Time, validation.Max(time.Unix(1502000000, 0))),
		)

		assert.NotNil(t, err)
		assert.EqualError(t, err, "Float: must be no greater than 500; Int: must be no greater than 500; Uint: must be no greater than 500.")
		assert.Contains(t, errTime.Error(), "Time: must be no greater than")
	})
}

func Test_GreaterOrEqual(t *testing.T) {
	type Struct struct {
		Float float64
		Int   int64
		Uint  uint64
		Time  time.Time
	}

	t.Run("ok - value valid", func(t *testing.T) {
		x := Struct{
			Float: 512.22,
			Int:   512,
			Uint:  512,
			Time:  time.Unix(1512000000, 0),
		}
		err := validation.ValidateStruct(&x,
			validation.Field(&x.Float, validation.GreaterOrEqual(500.0)),
			validation.Field(&x.Int, validation.GreaterOrEqual(500)),
			validation.Field(&x.Uint, validation.GreaterOrEqual(uint64(500))),
			validation.Field(&x.Time, validation.GreaterOrEqual(time.Unix(1502000000, 0))),
		)

		assert.Nil(t, err)
	})

	t.Run("ok - value equal number", func(t *testing.T) {
		x := Struct{
			Float: 512.22,
			Int:   512,
			Uint:  512,
			Time:  time.Unix(1512000000, 0),
		}
		err := validation.ValidateStruct(&x,
			validation.Field(&x.Float, validation.GreaterOrEqual(512.22)),
			validation.Field(&x.Int, validation.GreaterOrEqual(512)),
			validation.Field(&x.Uint, validation.GreaterOrEqual(uint64(512))),
			validation.Field(&x.Time, validation.GreaterOrEqual(time.Unix(1512000000, 0))),
		)

		assert.Nil(t, err)
	})

	t.Run("err - return error if value less than number", func(t *testing.T) {
		x := Struct{
			Float: 512.22,
			Int:   512,
			Uint:  512,
			Time:  time.Unix(1512000000, 0),
		}
		err := validation.ValidateStruct(&x,
			validation.Field(&x.Float, validation.GreaterOrEqual(1500.0)),
			validation.Field(&x.Int, validation.GreaterOrEqual(1500)),
			validation.Field(&x.Uint, validation.GreaterOrEqual(uint64(1500))),
		)
		errTime := validation.ValidateStruct(&x,
			validation.Field(&x.Time, validation.GreaterOrEqual(time.Unix(1602000000, 0))),
		)

		assert.NotNil(t, err)
		assert.EqualError(t, err, "Float: must be no less than 1500; Int: must be no less than 1500; Uint: must be no less than 1500.")
		assert.Contains(t, errTime.Error(), "Time: must be no less than") // it shows local timezone on the error message
	})
}

func Test_LessOrEqual(t *testing.T) {
	type Struct struct {
		Float float64
		Int   int64
		Uint  uint64
		Time  time.Time
	}

	t.Run("ok - value valid", func(t *testing.T) {
		x := Struct{
			Float: 512.22,
			Int:   512,
			Uint:  512,
			Time:  time.Unix(1512000000, 0),
		}
		err := validation.ValidateStruct(&x,
			validation.Field(&x.Float, validation.LessOrEqual(1500.0)),
			validation.Field(&x.Int, validation.LessOrEqual(1500)),
			validation.Field(&x.Uint, validation.LessOrEqual(uint64(1500))),
			validation.Field(&x.Time, validation.LessOrEqual(time.Unix(1602000000, 0))),
		)

		assert.Nil(t, err)
	})

	t.Run("ok - value equal number", func(t *testing.T) {
		x := Struct{
			Float: 512.22,
			Int:   512,
			Uint:  512,
			Time:  time.Unix(1512000000, 0),
		}
		err := validation.ValidateStruct(&x,
			validation.Field(&x.Float, validation.LessOrEqual(512.22)),
			validation.Field(&x.Int, validation.LessOrEqual(512)),
			validation.Field(&x.Uint, validation.LessOrEqual(uint64(512))),
			validation.Field(&x.Time, validation.LessOrEqual(time.Unix(1512000000, 0))),
		)

		assert.Nil(t, err)
	})

	t.Run("err - return error if value is greater than number", func(t *testing.T) {
		x := Struct{
			Float: 512.22,
			Int:   512,
			Uint:  512,
			Time:  time.Unix(1512000000, 0),
		}
		err := validation.ValidateStruct(&x,
			validation.Field(&x.Float, validation.LessOrEqual(500.0)),
			validation.Field(&x.Int, validation.LessOrEqual(500)),
			validation.Field(&x.Uint, validation.LessOrEqual(uint64(500))),
		)
		errTime := validation.ValidateStruct(&x,
			validation.Field(&x.Time, validation.LessOrEqual(time.Unix(1502000000, 0))),
		)

		assert.NotNil(t, err)
		assert.EqualError(t, err, "Float: must be no greater than 500; Int: must be no greater than 500; Uint: must be no greater than 500.")
		assert.Contains(t, errTime.Error(), "Time: must be no greater than")
	})
}
