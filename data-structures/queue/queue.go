package queue

type (
	Queue struct {
		len  int
		head *Node
		tail *Node
	}

	Node struct {
		val  int
		next *Node
	}

	Interface interface {
		Enqueue(val int)
		Dequeue() *Node
	}
)

func NewNode(val int) *Node {
	n := Node{
		val: val,
	}

	return &n
}

func New(val int) *Queue {
	n := NewNode(val)
	q := Queue{
		len:  1,
		head: n,
		tail: n,
	}

	return &q
}

func (q *Queue) Enqueue(val int) {
	n := NewNode(val)
	if q.len == 0 {
		q.head = n
		q.tail = n
		q.len++
		return
	}

	n.next = q.tail
	q.tail = n
	q.len++
}

func (q *Queue) Dequeue() *Node {
	if q.head == nil {
		return q.head
	}

	tmp := q.head
	q.head = q.head.next
	q.len--

	if q.len == 1 {
		q.tail = q.head
	}

	return tmp
}
