package ctrie

import "sync/atomic"

var gen uint64

type Ctrie[K comparable, V any] struct {
	inode    *inode
	hashfn   func(k any) uint64
	readonly bool
}

func generation() uint64 {
	return atomic.AddUint64(&gen, 1)
}
