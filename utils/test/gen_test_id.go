package test

import "gitlab.com/trivery-id/skadi/utils/random"

const testIDLen = 10

func GenerateTestID() string {
	return random.RandStringRunes(testIDLen)
}
