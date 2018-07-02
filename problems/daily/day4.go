package daily

// This problem was asked by Stripe.
// Given an array of integers, find the first missing positive integer in linear time and constant space. In other words, find the lowest positive integer that does not exist in the array. The array can contain duplicates and negative numbers as well.
// For example, the input [3, 4, -1, 1] should give 2. The input [1, 2, 0] should give 3.
// You can modify the input array in-place.

// lowestNotIncluded will return the lowest integer that is not included in the array.
// This function set min to 1 then iterates over all integers and adds them to a hash map.
// If the current integer is equal to min, it looks for the next smallest integer which is not in the hash map (i.e. hasn't been visited).
// After iterating over all in[i], min is returned.
//
// Another solution includes either sorting the input array, or building a Min Heap.
// The trade-off is time over space because those solutions take O(nlog(n)) runtime, but O(1) space.
//
// runtime: O(n)
// space:   O(n) because of the hashmap
func lowestNotIncluded(in []int) int {
	n := len(in)
	if n == 0 {
		return 0
	}

	hm := map[int]bool{}
	min := 1

	for i := 0; i < n; i++ {
		hm[in[i]] = true
		if in[i] == min {
			for hm[min] {
				min++
			}
		}
	}

	return min
}
