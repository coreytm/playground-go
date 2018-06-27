package bst

type Node struct {
	val   int
	left  *Node
	right *Node
}

type BinarySearchTree struct {
	height int
	len    int
	root   *Node
}

func NewNode(val int) *Node {
	n := Node{
		val: val,
	}

	return &n
}

func NewBST(val int) *BinarySearchTree {
	n := NewNode(val)
	bst := BinarySearchTree{
		root:   n,
		len:    1,
		height: 1,
	}

	return &bst
}
