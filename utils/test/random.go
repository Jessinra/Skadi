package test

import (
	"fmt"
	"time"
)

// Random generate random string with prefix.
func Random(prefix string) string {
	return fmt.Sprintf("%s-%d", prefix, time.Now().UnixNano())
}

func RandomInt64() int64 {
	return time.Now().UnixNano() % 7103 // nolint // random prime number
}
