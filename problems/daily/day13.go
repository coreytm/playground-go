package daily

// This problem was asked by Amazon.
// Given an integer k and a string s, find the length of the longest substring that contains at most k distinct characters.
// For example, given s = "abcba" and k = 2, the longest substring with k distinct characters is "bcb".

// maxSubstringLessThanKDistinct will return the max length of the substring in 'in' with at most 'k' distinct chars.
//
// runtime: O(n) because will iterate at most 2n to move p1 and p2
// space:   O(n) because we use a hash map to count of chars in substring
func maxSubstringLessThanKDistinct(in string, k int) (int, string) {
	n := len(in)
	if n == 0 || k <= 0 {
		return 0, ""
	}

	// use a map to keep count of letters in substring
	m := map[byte]int{}

	numDistinct, max, p1, p2 := 0, 0, 0, 0
	for i := 0; i < n; i++ {
		p2 = i

		// keep count of each char in substring
		if v, ok := m[in[p2]]; !ok || v == 0 {
			numDistinct++
			numP2 := m[in[p2]]
			m[in[p2]] = numP2 + 1
		}

		// remove the first character from substring until there are only 'k' distinct chars
		for numDistinct > k {
			numP1 := m[in[p1]] - 1
			m[in[p1]] = numP1
			if numP1 == 0 {
				numDistinct--
			}
			p1++
		}

		// replace max if the new substring is longer
		newMax := p2 - p1 + 1
		if newMax > max {
			max = newMax
		}
	}

	return max, in[p1 : p2+1]
}
