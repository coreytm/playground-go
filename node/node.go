package node

type Node struct {
	Value int
	Next  *Node
	Prev  *Node
}

func New(val int) *Node {
	n := Node{
		Value: val,
	}

	return &n
}
