package random_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	. "gitlab.com/trivery-id/skadi/utils/random"
)

func TestRandStringRunes(t *testing.T) {
	type args struct {
		n    int
		opts []Option
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "ok - generate 40 random string using default runes",
			args: args{
				n: 40,
			},
			want: "[abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789]+",
		},
		{
			name: "ok - generate 40 random string using spesific runes (1)",
			args: args{
				n: 40,
				opts: []Option{
					WithRunes([]rune("abc")),
				},
			},
			want: "[abc]+",
		},
		{
			name: "ok - generate 40 random string using spesific runes (2)",
			args: args{
				n: 40,
				opts: []Option{
					WithRunes([]rune("abc")),
				},
			},
			want: "[^defghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789]+",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RandStringRunes(tt.args.n, tt.args.opts...)
			assert.Len(t, got, tt.args.n)
			assert.Regexp(t, tt.want, got)
		})
	}
}

func TestGenerate10000UniqueString(t *testing.T) {
	strings := make(map[string]bool)

	for i := 0; i < 10000; i++ {
		got := RandStringRunes(40)
		if _, ok := strings[got]; ok {
			assert.False(t, ok)
		} else {
			strings[got] = true
		}
	}
	assert.Len(t, strings, 10000)
}
