package validation_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gitlab.com/trivery-id/skadi/utils/validation"
)

func Test_IsBool(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name       string
		args       args
		wantErr    bool
		wantErrMsg string
	}{
		{
			name: "ok - is boolean",
			args: args{
				value: true,
			},
			wantErr: false,
		},
		{
			name: "ok - is boolean",
			args: args{
				value: false,
			},
			wantErr: false,
		},
		{
			name: "ok - is string",
			args: args{
				value: "true",
			},
			wantErr:    true,
			wantErrMsg: "value: must be boolean.",
		},
		{
			name: "err - is flaot",
			args: args{
				value: 1,
			},
			wantErr:    true,
			wantErrMsg: "value: must be boolean.",
		},
		{
			name:       "err - is nil",
			args:       args{},
			wantErr:    true,
			wantErrMsg: "value: must be boolean.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validation.ValidateStruct(&tt.args,
				validation.Field(&tt.args.value, validation.IsBool),
			)

			assert.Equal(t, tt.wantErr, (err != nil))
			if tt.wantErr {
				assert.EqualError(t, err, tt.wantErrMsg)
			}
		})
	}
}

func Test_IsFloat(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name       string
		args       args
		wantErr    bool
		wantErrMsg string
	}{
		{
			name: "ok - is float32",
			args: args{
				value: float32(1.23),
			},
			wantErr: false,
		},
		{
			name: "ok - is float64",
			args: args{
				value: float64(1.23),
			},
			wantErr: false,
		},
		{
			name: "ok - is float",
			args: args{
				value: 1.23,
			},
			wantErr: false,
		},
		{
			name: "err - empty",
			args: args{
				value: nil,
			},
			wantErr:    true,
			wantErrMsg: "value: must be float32 or float64.",
		},
		{
			name: "err - float in string",
			args: args{
				value: "1.23",
			},
			wantErr:    true,
			wantErrMsg: "value: must be float32 or float64.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validation.ValidateStruct(&tt.args,
				validation.Field(&tt.args.value, validation.IsFloat),
			)

			assert.Equal(t, tt.wantErr, (err != nil))
			if tt.wantErr {
				assert.EqualError(t, err, tt.wantErrMsg)
			}
		})
	}
}

func Test_IsFloatLatitudeLongitude(t *testing.T) {
	type args struct {
		latitude  interface{}
		longitude interface{}
	}
	tests := []struct {
		name       string
		args       args
		wantErr    bool
		wantErrMsg string
	}{
		{
			name: "ok - is float coordinates",
			args: args{
				latitude:  35.71148215488324,
				longitude: 139.81061120566628,
			},
			wantErr: false,
		},
		{
			name: "ok - is negative float coordinates",
			args: args{
				latitude:  -35.71148215488324,
				longitude: -139.81061120566628,
			},
			wantErr: false,
		},
		{
			name: "ok - zero coordinates",
			args: args{
				latitude:  0.0,
				longitude: 0.0,
			},
			wantErr: false,
		},
		{
			name: "err - invalid value",
			args: args{
				latitude:  333335.71148215488324,
				longitude: 333339.81061120566628,
			},
			wantErr:    true,
			wantErrMsg: "latitude: must be a valid latitude; longitude: must be a valid longitude.",
		},
		{
			name: "err - empty",
			args: args{
				latitude:  nil,
				longitude: nil,
			},
			wantErr:    true,
			wantErrMsg: "latitude: must be float32 or float64; longitude: must be float32 or float64.",
		},
		{
			name: "err - coordinate in string",
			args: args{
				latitude:  "35.71148215488324",
				longitude: "139.81061120566628",
			},
			wantErr:    true,
			wantErrMsg: "latitude: must be float32 or float64; longitude: must be float32 or float64.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validation.ValidateStruct(&tt.args,
				validation.Field(&tt.args.latitude, validation.IsFloatLatitude),
				validation.Field(&tt.args.longitude, validation.IsFloatLongitude),
			)

			assert.Equal(t, tt.wantErr, (err != nil))
			if tt.wantErr {
				assert.EqualError(t, err, tt.wantErrMsg)
			}
		})
	}
}

func Test_IsInt(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name       string
		args       args
		wantErr    bool
		wantErrMsg string
	}{
		{
			name: "ok - is int",
			args: args{
				value: int(123),
			},
			wantErr: false,
		},
		{
			name: "ok - is int8",
			args: args{
				value: int8(123),
			},
			wantErr: false,
		},
		{
			name: "ok - is int16",
			args: args{
				value: int16(123),
			},
			wantErr: false,
		},
		{
			name: "ok - is int32",
			args: args{
				value: int32(123),
			},
			wantErr: false,
		},
		{
			name: "ok - is int64",
			args: args{
				value: int64(123),
			},
			wantErr: false,
		},
		{
			name: "err - empty",
			args: args{
				value: nil,
			},
			wantErr:    true,
			wantErrMsg: "value: must be an integer.",
		},
		{
			name: "err - string",
			args: args{
				value: "123",
			},
			wantErr:    true,
			wantErrMsg: "value: must be an integer.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validation.ValidateStruct(&tt.args,
				validation.Field(&tt.args.value, validation.IsInt),
			)

			assert.Equal(t, tt.wantErr, (err != nil))
			if tt.wantErr {
				assert.EqualError(t, err, tt.wantErrMsg)
			}
		})
	}
}

func Test_IsRoundNumber(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name       string
		args       args
		wantErr    bool
		wantErrMsg string
	}{
		{
			name: "ok - is int",
			args: args{
				value: int(123),
			},
			wantErr: false,
		},
		{
			name: "ok - is int8",
			args: args{
				value: int8(123),
			},
			wantErr: false,
		},
		{
			name: "ok - is int16",
			args: args{
				value: int16(123),
			},
			wantErr: false,
		},
		{
			name: "ok - is int32",
			args: args{
				value: int32(123),
			},
			wantErr: false,
		},
		{
			name: "ok - is int64",
			args: args{
				value: int64(123),
			},
			wantErr: false,
		},
		{
			name: "ok - is uint",
			args: args{
				value: uint(123),
			},
			wantErr: false,
		},
		{
			name: "ok - is uint8",
			args: args{
				value: uint8(123),
			},
			wantErr: false,
		},
		{
			name: "ok - is uint16",
			args: args{
				value: uint16(123),
			},
			wantErr: false,
		},
		{
			name: "ok - is uint32",
			args: args{
				value: uint32(123),
			},
			wantErr: false,
		},
		{
			name: "ok - is uint64",
			args: args{
				value: uint64(123),
			},
			wantErr: false,
		},
		{
			name: "ok - is no decimal float32",
			args: args{
				value: float32(123),
			},
			wantErr: false,
		},
		{
			name: "ok - is no decimal float64",
			args: args{
				value: float64(123),
			},
			wantErr: false,
		},
		{
			name: "err - is float64 with decimal",
			args: args{
				value: float64(123.123),
			},
			wantErr:    true,
			wantErrMsg: "value: must be an uint, int, or float without decimal points.",
		},
		{
			name: "err - empty",
			args: args{
				value: nil,
			},
			wantErr:    true,
			wantErrMsg: "value: must be an uint, int, or float without decimal points.",
		},
		{
			name: "err - string",
			args: args{
				value: "123",
			},
			wantErr:    true,
			wantErrMsg: "value: must be an uint, int, or float without decimal points.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validation.ValidateStruct(&tt.args,
				validation.Field(&tt.args.value, validation.IsRoundNumber),
			)

			assert.Equal(t, tt.wantErr, (err != nil))
			if tt.wantErr {
				assert.EqualError(t, err, tt.wantErrMsg)
			}
		})
	}
}

func Test_IsRFC3339(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name       string
		args       args
		wantErr    bool
		wantErrMsg string
	}{
		{
			name: "ok - valid format",
			args: args{
				value: time.Now().Format(time.RFC3339),
			},
			wantErr: false,
		},
		{
			name: "ok - valid format",
			args: args{
				value: "2021-07-02T02:00:59+00:00",
			},
			wantErr: false,
		},
		{
			name: "err - invalid format time.Time",
			args: args{
				value: time.Now(),
			},
			wantErr:    true,
			wantErrMsg: "value: must follow RFC3339 format.",
		},
		{
			name: "err - invalid format ISO8601",
			args: args{
				value: "2021-07-02T02:00:59+0000",
			},
			wantErr:    true,
			wantErrMsg: "value: must follow RFC3339 format.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validation.ValidateStruct(&tt.args,
				validation.Field(&tt.args.value, validation.IsRFC3339),
			)

			assert.Equal(t, tt.wantErr, (err != nil))
			if tt.wantErr {
				assert.EqualError(t, err, tt.wantErrMsg)
			}
		})
	}
}

func Test_IsDate(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name       string
		args       args
		wantErr    bool
		wantErrMsg string
	}{
		{
			name: "ok - valid format",
			args: args{
				value: time.Now().Format("2006-01-02"),
			},
			wantErr: false,
		},
		{
			name: "ok - valid format",
			args: args{
				value: "2021-07-02",
			},
			wantErr: false,
		},
		{
			name: "err - invalid format time.Time",
			args: args{
				value: time.Now().Format(time.RFC3339),
			},
			wantErr:    true,
			wantErrMsg: "value: must follow ISO8601 format.",
		},
		{
			name: "err - invalid format ISO8601",
			args: args{
				value: "2021-07-02T02:00:59+0000",
			},
			wantErr:    true,
			wantErrMsg: "value: must follow ISO8601 format.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validation.ValidateStruct(&tt.args,
				validation.Field(&tt.args.value, validation.IsDate),
			)

			assert.Equal(t, tt.wantErr, (err != nil))
			if tt.wantErr {
				assert.EqualError(t, err, tt.wantErrMsg)
			}
		})
	}
}

func Test_IsPositiveOrZero(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name       string
		args       args
		wantErr    bool
		wantErrMsg string
	}{
		{
			name: "ok - int",
			args: args{
				value: int(123),
			},
			wantErr: false,
		},
		{
			name: "ok - int8",
			args: args{
				value: int8(123),
			},
			wantErr: false,
		},
		{
			name: "ok - int16",
			args: args{
				value: int16(123),
			},
			wantErr: false,
		},
		{
			name: "ok - int32",
			args: args{
				value: int32(123),
			},
			wantErr: false,
		},
		{
			name: "ok - int64",
			args: args{
				value: int64(123),
			},
			wantErr: false,
		},
		{
			name: "ok - uint",
			args: args{
				value: uint(123),
			},
			wantErr: false,
		},
		{
			name: "ok - uint8",
			args: args{
				value: uint8(123),
			},
			wantErr: false,
		},
		{
			name: "ok - uint16",
			args: args{
				value: uint16(123),
			},
			wantErr: false,
		},
		{
			name: "ok - uint32",
			args: args{
				value: uint32(123),
			},
			wantErr: false,
		},
		{
			name: "ok - uint64",
			args: args{
				value: uint64(123),
			},
			wantErr: false,
		},
		{
			name: "ok - float32",
			args: args{
				value: float32(1.23),
			},
			wantErr: false,
		},
		{
			name: "ok - float64",
			args: args{
				value: float64(1.23),
			},
			wantErr: false,
		},
		{
			name: "err - negative int",
			args: args{
				value: int(-123),
			},
			wantErr:    true,
			wantErrMsg: "value: must be positive or zero value.",
		},
		{
			name: "err - negative int8",
			args: args{
				value: int8(-123),
			},
			wantErr:    true,
			wantErrMsg: "value: must be positive or zero value.",
		},
		{
			name: "err - negative int16",
			args: args{
				value: int16(-123),
			},
			wantErr:    true,
			wantErrMsg: "value: must be positive or zero value.",
		},
		{
			name: "err - negative int32",
			args: args{
				value: int32(-123),
			},
			wantErr:    true,
			wantErrMsg: "value: must be positive or zero value.",
		},
		{
			name: "err - negative int64",
			args: args{
				value: int64(-123),
			},
			wantErr:    true,
			wantErrMsg: "value: must be positive or zero value.",
		},
		{
			name: "err - negative float32",
			args: args{
				value: float32(-1.23),
			},
			wantErr:    true,
			wantErrMsg: "value: must be positive or zero value.",
		},
		{
			name: "err - negative float64",
			args: args{
				value: float64(-1.23),
			},
			wantErr:    true,
			wantErrMsg: "value: must be positive or zero value.",
		},
		{
			name: "err - empty",
			args: args{
				value: nil,
			},
			wantErr:    true,
			wantErrMsg: "value: must be a number.",
		},
		{
			name: "err - string",
			args: args{
				value: "123",
			},
			wantErr:    true,
			wantErrMsg: "value: must be a number.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validation.ValidateStruct(&tt.args,
				validation.Field(&tt.args.value, validation.IsPositiveOrZero),
			)

			assert.Equal(t, tt.wantErr, (err != nil))
			if tt.wantErr {
				assert.EqualError(t, err, tt.wantErrMsg)
			}
		})
	}
}
