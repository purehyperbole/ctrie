package ctrie

// idirection node
type inode struct {
	main       *mnode
	generation uint64
}
