package ctrie

// list node
type lnode[K comparable, V any] struct {
	sn   snode[K, V]
	next *lnode[K, V]
}
