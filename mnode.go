package ctrie

import "unsafe"

type mtype byte

const (
	mtypec mtype = 0
	mtypet mtype = 1
	mtypel mtype = 2
)

type mnode struct {
	typ  mtype
	ptr  unsafe.Pointer
	prev *mnode
}
