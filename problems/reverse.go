package problems

type (
	node struct {
		val  int
		next *node
	}
)

func (n *node) reverse() *node {
	cur := n
	var prev *node
	var next *node

	for cur != nil {
		next = cur.next
		cur.next = prev
		prev = cur
		cur = next
	}

	return prev
}
