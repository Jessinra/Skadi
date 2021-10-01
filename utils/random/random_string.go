package random

import (
	"crypto/rand"
	"math/big"
)

func RandStringRunes(n int, opts ...Option) string {
	option := parseOptions(opts...)

	b := make([]rune, n)
	for i := range b {
		randomNum, _ := rand.Int(rand.Reader, big.NewInt(int64(len(option.runes))))
		b[i] = option.runes[randomNum.Int64()]
	}

	return string(b)
}
