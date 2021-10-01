package masking_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/trivery-id/skadi/utils/masking"
)

func TestEmail(t *testing.T) {
	tests := []struct {
		name  string
		email string
		want  string
	}{
		{
			name:  "ok - common masked email",
			email: "totos21s@gmail.com",
			want:  "t*****1s@gmail.com",
		},
		{
			name:  "ok - invalid email without domain, masked anyway",
			email: "totos21sgmail",
			want:  "to********ail",
		},
		{
			name:  "ok - single char, masked all",
			email: "a@gmail.com",
			want:  "*@gmail.com",
		},
		{
			name:  "ok - 2 chars, masked all but last char",
			email: "ab@gmail.com",
			want:  "*b@gmail.com",
		},
		{
			name:  "ok - 3 chars, masked all but last char",
			email: "abc@gmail.com",
			want:  "**c@gmail.com",
		},
		{
			name:  "ok - 4 chars, masked all but last char",
			email: "abcd@gmail.com",
			want:  "***d@gmail.com",
		},
		{
			name:  "ok - 5 char and above, masked middle 60%",
			email: "abcde@gmail.com",
			want:  "a***e@gmail.com",
		},
		{
			name:  "ok - invalid empty address, no changes",
			email: "@gmail.com",
			want:  "@gmail.com",
		},
		{
			name:  "ok - empty",
			email: "",
			want:  "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := masking.Email(tt.email)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestPhone(t *testing.T) {
	tests := []struct {
		name  string
		phone string
		want  string
	}{
		{
			name:  "ok - normal phone number, get first 3 and last 3",
			phone: "+6287800990099",
			want:  "+62********099",
		},
		{
			name:  "ok - normal phone number, get first 3 and last 3",
			phone: "+187800990099",
			want:  "+18*******099",
		},
		{
			name:  "ok - normal phone number, get first 3 and last 3",
			phone: "087800990099",
			want:  "087******099",
		},
		{
			name:  "ok - less than 6 char phone number (invalid), return all",
			phone: "14022",
			want:  "14022",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := masking.Phone(tt.phone)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestName(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{
			name: "ok - common masked name",
			in:   "Richard James Wayne",
			want: "Ri***** Ja*** Wa***",
		},
		{
			name: "ok - common masked name",
			in:   "Leona Hearthfelt",
			want: "Le*** Hear******",
		},
		{
			name: "ok - one word name",
			in:   "Diana",
			want: "Di***",
		},
		{
			name: "ok - single char, masked all",
			in:   "A",
			want: "*",
		},
		{
			name: "ok - 2 chars, masked all",
			in:   "Po",
			want: "**",
		},
		{
			name: "ok - 3 chars, masked all but first char",
			in:   "Rio",
			want: "R**",
		},
		{
			name: "ok - 4 chars, masked all but first char",
			in:   "Budi",
			want: "B***",
		},
		{
			name: "ok - 5 char and above, masked right 60%",
			in:   "Diana",
			want: "Di***",
		},
		{
			name: "ok - empty",
			in:   "",
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := masking.Name(tt.in)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestLeftRightCenter(t *testing.T) {
	type args struct {
		s           string
		maskPercent float64
	}
	tests := []struct {
		name       string
		args       args
		wantLeft   string
		wantRight  string
		wantCenter string
	}{
		{
			name: "ok - valid string (len 5)",
			args: args{
				s:           "abcde",
				maskPercent: 0.6,
			},
			wantLeft:   "***de",
			wantRight:  "ab***",
			wantCenter: "a***e",
		},
		{
			name: "ok - valid string (len 15)",
			args: args{
				s:           "abcdefghijklmno",
				maskPercent: 1.0 / 3,
			},
			wantLeft:   "*****fghijklmno",
			wantRight:  "abcdefghij*****",
			wantCenter: "abcde*****klmno",
		},
		{
			name: "ok - 1 char low percentage",
			args: args{
				s:           "a",
				maskPercent: 0.2,
			},
			wantLeft:   "a",
			wantRight:  "a",
			wantCenter: "a",
		},
		{
			name: "ok - 1 char high percentage",
			args: args{
				s:           "a",
				maskPercent: 0.8,
			},
			wantLeft:   "*",
			wantRight:  "*",
			wantCenter: "*",
		},
		{
			name: "ok - mask all",
			args: args{
				s:           "abcdef",
				maskPercent: 1,
			},
			wantLeft:   "******",
			wantRight:  "******",
			wantCenter: "******",
		},
		{
			name: "ok - mask none",
			args: args{
				s:           "abcdef",
				maskPercent: 0,
			},
			wantLeft:   "abcdef",
			wantRight:  "abcdef",
			wantCenter: "abcdef",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLeft := masking.Left(tt.args.s, tt.args.maskPercent)
			assert.Equal(t, tt.wantLeft, gotLeft)

			gotRight := masking.Right(tt.args.s, tt.args.maskPercent)
			assert.Equal(t, tt.wantRight, gotRight)

			gotCenter := masking.Center(tt.args.s, tt.args.maskPercent)
			assert.Equal(t, tt.wantCenter, gotCenter)
		})
	}
}
