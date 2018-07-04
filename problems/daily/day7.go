package daily

// This problem was asked by Facebook.
// Given the mapping a = 1, b = 2, ... z = 26, and an encoded message, count the number of ways it can be decoded.
// For example, the message '111' would give 3, since it could be decoded as 'aaa', 'ka', and 'ak'.
// You can assume that the messages are decodable. For example, '001' is not allowed.

// numDecodeRecursive will return the number of decoded permutations.
// This function will consider the last single digit int, and last double digit int in the provided array and recurse over the sub-array accordingly.
func numDecodeRecursive(in []int) int {
	n := len(in)

	if n <= 1 {
		return 1
	}

	numDecode := 0
	a := in[n-1]
	if a > 0 && a < 10 {
		numDecode += numDecodeRecursive(in[:n])
	}

	b := (in[n-2] * 10) + in[n-1]
	if b >= 10 && b <= 26 {
		numDecode += numDecodeRecursive(in[:n-1])
	}

	return numDecode
}
