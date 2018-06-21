package stack

type (
	Stack struct {
		len  int
		head *Node
		tail *Node
	}

	Node struct {
		val  int
		next *Node
	}

	Interface interface {
		Push(val int)
		Pop() *Node
		Peek() *Node
	}
)

func NewNode(val int) *Node {
	n := Node{
		val: val,
	}

	return &n
}

func New(val int) *Stack {
	n := Node{
		val: val,
	}

	s := Stack{
		len:  1,
		head: &n,
		tail: &n,
	}

	return &s
}

func (s *Stack) Push(val int) {
	n := NewNode(val)

	if s.len == 0 {
		s.head = n
		s.tail = n
		s.len++
		return
	}

	n.next = s.head
	if s.len == 1 {
		s.head.next = s.tail
	}

	s.head = n
	s.len++
}

func (s *Stack) Pop() *Node {
	if s.len == 0 {
		return nil
	}

	if s.len == 1 {
		s.tail = nil
	}

	tmp := s.head
	s.head = tmp.next
	s.len--
	return tmp
}

func (s *Stack) Peek() *Node {
	return s.head
}
