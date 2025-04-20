package ctrie

import (
	"sync/atomic"
	"unsafe"
)

// Generation-compare-and-swap, or GCAS. This procedure has semantics
// similar to that of the RDCSS, but it does not create the intermediate
// object except in the case of failures that occur due to the snapshot
// being taken
func gcas[K Key, V any](trie *Ctrie[K, V], currentNode *inode, oldNode, newNode *mnode) bool {
	atomic.StorePointer(
		(*unsafe.Pointer)(unsafe.Pointer(&newNode.prev)),
		unsafe.Pointer(oldNode),
	)

	if !atomic.CompareAndSwapPointer(
		(*unsafe.Pointer)(unsafe.Pointer(&currentNode.main)),
		unsafe.Pointer(oldNode),
		unsafe.Pointer(newNode),
	) {
		return false
	}

	gcasCommit()

	return atomic.LoadPointer(
		(*unsafe.Pointer)(unsafe.Pointer(&newNode.prev)),
	) == nil
}

func gcasCommit[K Key, V any](trie *Ctrie[K, V], currentNode *inode, mainNode *mnode) unsafe.Pointer {
	previous := atomic.LoadPointer(
		(*unsafe.Pointer)(unsafe.Pointer(&mainNode.prev)),
	)

	// TODO finish

	return nil
}

func gcasRead[K Key, V any](trie *Ctrie[K, V], currentNode *inode) unsafe.Pointer {
	main := atomic.LoadPointer(
		(*unsafe.Pointer)(unsafe.Pointer(&currentNode.main)),
	)

	previous := atomic.LoadPointer(
		(*unsafe.Pointer)(unsafe.Pointer(&(*mnode)(main).prev)),
	)

	if previous == nil {
		return main
	}

	return gcasCommit(trie, currentNode, (*mnode)(main))
}
