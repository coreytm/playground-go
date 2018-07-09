package leetcode

// 779. K-th Symbol in Grammar
// On the first row, we write a 0. Now in every subsequent row, we look at the previous row and replace each occurrence of 0 with 01, and each occurrence of 1 with 10.
//
// Given row N and index K, return the K-th indexed symbol in row N. (The values of K are 1-indexed.) (1 indexed).
//
// Examples:
// Input: N = 1, K = 1
// Output: 0
//
// Input: N = 2, K = 1
// Output: 0
//
// Input: N = 2, K = 2
// Output: 1
//
// Input: N = 4, K = 5
// Output: 1
//
// Explanation:
// row 1: 0
// row 2: 01
// row 3: 0110
// row 4: 01101001
//
// Note:
// N will be an integer in the range [1, 30].
// K will be an integer in the range [1, 2^(N-1)].

// kthSymbol will return the byte value 0 or 1 of the kth index on row n.
// runtime: O(n)
// space: O(n)
func kthSymbol(n, k int) byte {
	// deal with base cases since can't do byte lookup on perms with size less than 4 bits
	if n < 2 {
		if k == 0 {
			return byte(0)
		}

		return byte(1)
	}

	// start with the first full byte perm
	b := []byte{6}
	// already dealt with base cases, so only need to iterate to n-2
	for i := 1; i < n-2; i++ {
		mid := len(b) / 2
		front := b[:mid]
		back := b[mid:]
		b = append(b, append(back, front...)...)
	}

	bytePos := (k - 1) / 4
	bitPos := (k - 1) % 4
	bitLookup := []int{3, 2, 1, 0}

	return b[bytePos] & byte(2^bitLookup[bitPos])
}
