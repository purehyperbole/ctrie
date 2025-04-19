package ctrie

// value node
type snode[K comparable, V any] struct {
	key   K
	value V
}
