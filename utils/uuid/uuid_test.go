package uuid_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	. "gitlab.com/trivery-id/skadi/utils/uuid"
)

func Test_NewUUID(t *testing.T) {
	t.Run("ok - generated UUID", func(t *testing.T) {
		got := NewUUID()
		assert.NotEmpty(t, got)
		assert.Len(t, got, 36)
	})

	t.Run("ok - no repeating UUID in 1 mil attempt", func(t *testing.T) {
		uuids := map[string]bool{}

		for i := 0; i < 1000000; i++ {
			got := NewUUID()
			if uuids[got] {
				t.Fatalf("found repeating UUID at %d attempt", i)
			}

			uuids[got] = true
		}
	})
}
