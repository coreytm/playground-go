package daily

// This problem was asked by Google.
// Given a list of numbers and a number k, return whether any two numbers from the list add up to k.
// For example, given [10, 15, 3, 7] and k of 17, return true since 10 + 7 is 17.
// Bonus: Can you do this in one pass?

// addsToK returns whether or not there are two integers in 'in' that add to 'k'.
// runtime: O(n)
// space:   O(n) ...is there a way to do this in O(1) space?
func addsToK(in []int, k int) bool {
	n := len(in)
	if n < 2 {
		return false
	}

	hm := map[int]bool{}
	for num := range in {
		if hm[k-num] {
			return true
		}

		hm[num] = true
	}

	return false
}
