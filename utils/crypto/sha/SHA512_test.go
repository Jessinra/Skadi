package sha_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/trivery-id/skadi/utils/crypto/sha"
)

func TestHash512(t *testing.T) {
	const expectedLen = 88

	tests := []struct {
		name string
		str  string
		want string
	}{
		{
			name: "ok - hashed",
			str:  "hello world",
			want: "MJ7MSJwS1utMxA9QyQLytNDtd-5RGnx6m808qG1M2G-YndNbxf9JlnDaNCVbRbDP2DDoH2Bdz33FVC6TrpzXbw==",
		},
		{
			name: "ok - hashed",
			str:  "konichiwa sekai",
			want: "0XR6wgBE29EelVACOQ67AKvWP50iD6MLZvv_ZfX6iBCyTO4tTx9A02sdknYOb-TK-yNCSikrGrO7nR61Afm3nA==",
		},
		{
			name: "ok - empty input, also hashed",
			str:  "",
			want: "z4PhNX7vuL3xVChQ1m2AB9Yg5AULVxXcg_SpIdNs6c5H0NE8XYXysP-DGNKHfuwvY7kxvUdBeoGlODJ6-SfaPg==",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sha.Hash512(tt.str)
			assert.Equal(t, tt.want, got)
			assert.Len(t, got, expectedLen)
		})
	}
}
