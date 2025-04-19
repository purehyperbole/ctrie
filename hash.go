package ctrie

import "unsafe"

//go:noescape
//go:linkname memhash runtime.memhash
func memhash(p unsafe.Pointer, h, s uintptr) uintptr

type Key interface {
	[]byte | string
}

// NOTE dynamically returning this as a lambda function
// for each key type as a one time setup call was dramatically
// slower (2x) than using a single static function. May be due to
// having to dereference the returned function or it cannot be inlined?
func hashfn[K Key](k K) uint64 {
	h := (*header)(unsafe.Pointer(&k))
	return uint64(memhash(h.ptr, 0, uintptr(h.len)))
}

type header struct {
	ptr unsafe.Pointer
	len int
}
