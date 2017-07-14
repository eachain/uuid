package uuid

import (
	"crypto/rand"
	"sync"
)

const hex = "0123456789abcdefg"

var (
	pool sync.Pool
	idxs []int
)

func init() {
	pool.New = func() interface{} {
		return make([]byte, 36)
	}
	idxs = []int{8, 13, 18, 23}
}

// New returns a uuid string like:
// xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx.
func New() string {
	bs := pool.Get().([]byte)
	rand.Read(bs)
	for i, b := range bs {
		bs[i] = hex[b&15]
	}
	for _, idx := range idxs {
		bs[idx] = '-'
	}
	uuid := string(bs)
	pool.Put(bs)
	return uuid
}
