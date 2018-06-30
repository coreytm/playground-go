package daily

// This problem was asked by Uber.
// Given an array of integers, return a new array such that each element at index i of the new array is the product of all the numbers in the original array except the one at i.
// For example, if our input was [1, 2, 3, 4, 5], the expected output would be [120, 60, 40, 30, 24]. If our input was [3, 2, 1], the expected output would be [2, 3, 6].
// Follow-up: what if you can't use division?

// productExcludingI returns a slice of integers which contains the product of all the numbers in 'in' excluding 'i' at in[i]
// runtime: O(n)
// space: O(n)
func productExcludingI(in []int) []int {
	n := len(in)
	if n < 2 {
		return []int{}
	}

	before := make([]int, n)
	after := make([]int, n)

	// pad with 1's because can't go left/right of first/last int
	before[0] = 1
	after[n-1] = 1

	for i := 1; i < n; i++ {
		before[i] = before[i-1] * in[i-1]
		after[n-1-i] = after[n-i] * in[n-i]
	}

	for i := 0; i < n; i++ {
		in[i] = before[i] * after[i]
	}

	return in
}
