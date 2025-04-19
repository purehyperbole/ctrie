package ctrie

import "unsafe"

// idirection node
type inode struct {
	mtype      mtype
	mnode      unsafe.Pointer
	generation uint64
}
