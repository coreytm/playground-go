package search

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

func (bst *BinarySearchTree) DepthFirstSearch(val int) *Node {
	if bst.len == 0 {
		return nil
	}

	return dfsRecurse(bst.root, val)
}

func dfsRecurse(root *Node, val int) *Node {
	if root == nil {
		return nil
	}

	for root.val != val {
		if root.val < val {
			return dfsRecurse(root.right, val)
		}

		return dfsRecurse(root.left, val)
	}

	return nil
}

// dfsIterative visits each node and adds to a stack.
// 1) Push right
// 2) Push left
// 3) Pop => if match return => else repeat 1 and 2
func dfsIterative(root *Node, val int) *Node {
	return nil
}

// bfsIterative visits each node and adds to a queue.
// 1) Enqueue left
// 2) Enqueue right
// 3) Dequeue => if match return => else repeat 1 and 2
func bfsIterative(root *Node, val int) *Node {
	return nil
}
