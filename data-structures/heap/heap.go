package heap

// Implemant heaps using Go's builtin interfaces for practice

// Heap is implemented using an array.
type Heap []int

// heap properties
// left := i * 2 + 1
// right := left + 1
// parent := i / 2 if i % 2 == 1
// parent := (i-1) / 2 if i % 2 == 0

// BuildMax will take an unordered array and turn it into a Max Heap.
// Since this function splits the dataset in half each iteration it
// performs in O(log(n)) runtime and O(1) space.
func BuildMax(in []int) Heap {
	h := Heap(in)

	pos := len(in) / 2
	for pos > 0 {
		h.MaxHeapifyDown(pos)
	}

	return nil
}

func (h Heap) Add(val int) {
	h = append(h, val)
	pos := len(h) - 1
	parentPos := (pos - 1) / 2
	if pos%2 == 1 {
		parentPos = pos / 2
	}

	if h[parentPos] > h[pos] {
		return
	}

	tmp := h[parentPos]
	h[parentPos] = h[pos]
	h[pos] = tmp
}

func (h Heap) PopMax() int {
	if len(h) == 0 {
		return -1
	}

	if len(h) == 1 {
		tmp := h[0]
		h = Heap{}
		return tmp
	}

	// swap first and last
	tmp := h[0]
	h[0] = h[len(h)-1]
	h = h[:len(h)-1]

	h.MaxHeapifyDown(0)

	return tmp
}

// MaxHeapifyDown will put the specified location in it's proper place in the Heap.
// It does this by swapping values of the provided position and the last position.
// Then it will bubble it up until it's in the correct spot.
func (h Heap) MaxHeapifyDown(pos int) {
	left := pos*2 + 1
	right := left + 1

	maxPos := pos
	if len(h) > left && h[left] > h[maxPos] {
		maxPos = left
	}
	if len(h) > right && h[right] > h[maxPos] {
		maxPos = right
	}

	if maxPos == pos {
		return
	}

	tmp := h[pos]
	h[pos] = h[maxPos]
	h[maxPos] = tmp

	h.MaxHeapifyDown(maxPos)
}

func (h Heap) MinHeapify(pos int) {}
