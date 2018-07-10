package daily

// This problem was asked by Google.
// Given two singly linked lists that intersect at some point, find the intersecting node. The lists are non-cyclical.
// For example, given A = 3 -> 7 -> 8 -> 10 and B = 99 -> 1 -> 8 -> 10, return the node with value 8.
// In this example, assume nodes with the same value are the exact same node objects.
// Do this in O(M + N) time (where M and N are the lengths of the lists) and constant space.

type linkedListNode struct {
	val  int
	next *linkedListNode
}

// findIntersectingNode returns the first intersecting node of the provided linked lists.
// If either of the linked lists are nil, returns nil.
//
// runtime: O(M + N) when M is the length of a and N is the length of b
// space:   O(n) because of the hasmap
func findIntersectingNode(a *linkedListNode, b *linkedListNode) *linkedListNode {
	if a == nil || b == nil {
		return nil
	}

	hm := map[int]*linkedListNode{}

	for a != nil {
		hm[a.val] = a
		a = a.next
	}

	for b != nil {
		if _, ok := hm[b.val]; ok {
			return b
		}

		b = b.next
	}

	return nil
}
