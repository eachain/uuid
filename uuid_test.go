package uuid_test

import (
	"testing"

	"github.com/eachain/uuid"
)

func TestNew(t *testing.T) {
	m := make(map[string]bool)
	for i := 0; i < 10000000; i++ {
		u := uuid.New()
		if m[u] {
			t.Fail()
		}
		m[u] = true
	}
}
