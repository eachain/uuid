package uuid

import (
	"testing"
)

func TestNew(t *testing.T) {
	m := make(map[string]bool)
	for i := 0; i < 10000000; i++ {
		u := New()
		if m[u] {
			t.Fail()
		}
		m[u] = true
	}
}
