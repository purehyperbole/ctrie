package ctrie

import (
	"sync/atomic"
	"unsafe"
)

// https://timharris.uk/papers/2002-disc.pdf

/*

def RDCSS(ov, ovmain, nv)
	r = READ(root)
	if r = ov ∧ GCAS_READ(ov.main) = ovmain {
		WRITE(root, nv)
		return >
	} else return ⊥


def snapshot()162
	r = RDCSS_READ()
	expmain = GCAS_READ(r)
	if RDCSS(r, expmain, new INode(expmain, new Gen))
		return new Ctrie {
		root = new INode(expmain, new Gen)
		readonly = ⊥
	}
	else return snapshot()
*/

func rdcss[K Key, V any](trie *Ctrie[K, V], currentNode *inode) {
	root := atomic.LoadPointer(
		(*unsafe.Pointer)(unsafe.Pointer(&trie.root)),
	)

}

func rcdssRead[K Key, V any](trie *Ctrie[K, V], currentNode *inode) {
}

func rcdssComplete[K Key, V any](trie *Ctrie[K, V], currentNode *inode) {

}
