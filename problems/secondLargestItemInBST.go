package problems

type (
	Node struct {
		val   int
		left  *Node
		right *Node
	}
)

func (n *Node) get