package daily

// This problem was asked by Airbnb.
// Given a list of integers, write a function that returns the largest sum of non-adjacent numbers. Numbers can be 0 or negative.
// For example, [2, 4, 6, 2, 5] should return 13, since we pick 2, 6, and 5. [5, 1, 1, 5] should return 10, since we pick 5 and 5.
// Follow-up: Can you do this in O(N) time and constant space?

// largestSumOfNonAdjacent uses a greedy approach to return the largest sum of non-adjacent integers.
// runtime: O(n) // iterates through input once and keeps track of largest so far
// space:   O(1) // utilizes existing memeory, however this contaminates the input
func largestSumOfNonAdjacent(in []int) int {
	n := len(in)

	if n == 0 {
		return 0
	}

	if n == 1 {
		return in[0]
	}

	if in[0] > in[1] {
		in[1] = in[0]
	}

	for i := 2; i < n; i++ {
		// If in[i] is larger than in[i-1] and (in[i] + in[i-2]) it will be unchanged.
		// This can occur if there are negative numbers in the list.
		if in[i-1] > in[i] {
			in[i] = in[i-1]
		}

		if in[i]+in[i-2] > in[i] {
			in[i] = in[i] + in[i-2]
		}
	}

	return in[n-1]
}
