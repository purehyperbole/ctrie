package ctrie

import "sync/atomic"

var gen uint64

type Ctrie[K Key, V any] struct {
	inode    *inode
	hashfn   func(k any) uint64
	readonly bool
}

func (c *Ctrie[K, V]) Lookup(key K) (V, bool) {

}

func ilookup[K Key, V any](k K, current, parent *inode, level uint) (V, bool, bool) {
	hash := hashfn(k)

}

// this *can* overflow, but is extremely unlikely to
// it will take decades to do so, even in a loop
// that does nothing but iterate.
func generation() uint64 {
	return atomic.AddUint64(&gen, 1)
}
