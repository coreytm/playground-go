package linkedList

type (
	LinkedList struct {
		len  int
		head *Node
		tail *Node
	}

	Node struct {
		val  int
		next *Node
		prev *Node
	}

	Interface interface {
		Add(val int)
		Delete(val int) *Node
	}
)

func NewNode(val int) *Node {
	n := Node{
		val: val,
	}

	return &n
}

func New(val int) *LinkedList {
	n := NewNode(val)
	l := LinkedList{
		len:  1,
		head: n,
		tail: n,
	}

	return &l
}

func (l *LinkedList) Add(val int) {
	n := NewNode(val)

	if l.len == 0 {
		l.head = n
		l.tail = n
		l.len++
		return
	}

	// a->b->c

	cur := l.head
	for cur.val <= n.val {
		cur = cur.next
	}

	if cur == l.head {
		n.next = l.head
		l.head.prev = n
		l.head = n
		l.len++
		return
	}

	if cur == nil {
		l.tail.next = n
		n.prev = l.tail
		l.len++
		return
	}

	n.next = cur
	n.prev = cur.prev
	cur.prev.next = n
	cur.next.prev = n
	l.len++
}

func (l *LinkedList) Delete(val int) *Node {
	if l.len == 0 {
		return nil
	}

	cur := l.head
	for cur != nil && cur.val != val {
		cur = cur.next
	}

	if cur == nil {
		return nil
	}

	if cur == l.head {
		l.head = l.head.next
		l.head.prev = nil
		l.len--
		return cur
	}

	if cur == l.tail {
		l.tail = cur.prev
		l.tail.next = nil
		l.len--
		return cur
	}

	cur.prev.next = cur.next
	cur.next.prev = cur.prev
	l.len--
	return cur
}
