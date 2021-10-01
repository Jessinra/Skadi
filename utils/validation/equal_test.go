package validation_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/trivery-id/skadi/utils/validation"
)

func TestEquals(t *testing.T) {
	type enum string
	const enumA enum = "ENUM_A"

	type something struct {
		value string
	}
	structA := something{value: "A"}
	structB := something{value: "B"}

	type args struct {
		value  interface{}
		target interface{}
	}
	tests := []struct {
		name       string
		args       args
		wantErr    bool
		wantErrMsg string
	}{
		{
			name: "ok - equal float",
			args: args{
				value:  1.23,
				target: 1.23,
			},
			wantErr: false,
		},
		{
			name: "err - not equal float",
			args: args{
				value:  1.00,
				target: 1.23,
			},
			wantErr:    true,
			wantErrMsg: "value: must equals '1.23' (float64).",
		},
		{
			name: "ok - equal string",
			args: args{
				value:  "abcd",
				target: "abcd",
			},
			wantErr: false,
		},
		{
			name: "err - equal string and enum (different type same value)",
			args: args{
				value:  "ENUM_A",
				target: enumA,
			},
			wantErr:    true,
			wantErrMsg: "value: must equals 'ENUM_A' (validation_test.enum).",
		},
		{
			name: "ok - equal string and enum (typecasted)",
			args: args{
				value:  enum("ENUM_A"),
				target: enumA,
			},
			wantErr: false,
		},
		{
			name: "err - not equal string",
			args: args{
				value:  "abcd",
				target: "hijk",
			},
			wantErr:    true,
			wantErrMsg: "value: must equals 'hijk' (string).",
		},
		{
			name: "ok - equal struct",
			args: args{
				value:  structA,
				target: structA,
			},
			wantErr: false,
		},
		{
			name: "ok - not equal struct",
			args: args{
				value:  structA,
				target: structB,
			},
			wantErr:    true,
			wantErrMsg: "value: must equals '{value:B}' (validation_test.something).",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validation.ValidateStruct(&tt.args,
				validation.Field(&tt.args.value, validation.Equals(tt.args.target)),
			)

			assert.Equal(t, tt.wantErr, (err != nil))
			if tt.wantErr {
				assert.EqualError(t, err, tt.wantErrMsg)
			}
		})
	}
}
