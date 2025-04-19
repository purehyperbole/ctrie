package ctrie

import "unsafe"

type cnode struct {
	bitmap uint64
	array  []unsafe.Pointer
}
