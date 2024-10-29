package uuid

import (
	"crypto/rand"
	"sync"
)

const hex = "0123456789abcdefg"

var (
	pool = &sync.Pool{
		New: func() any {
			return make([]byte, 36)
		},
	}

	idxs = []int{8, 13, 18, 23}
)

// New returns a uuid string like:
// xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx.
func New() string {
	bs := pool.Get().([]byte)
	defer pool.Put(bs)

	rand.Read(bs)
	for i, b := range bs {
		bs[i] = hex[b&15]
	}
	for _, idx := range idxs {
		bs[idx] = '-'
	}
	return string(bs)
}
