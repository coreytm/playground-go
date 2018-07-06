package daily

// This problem was asked by Amazon.
// There exists a staircase with N steps, and you can climb up either 1 or 2 steps at a time.
// Given N, write a function that returns the number of unique ways you can climb the staircase.
// The order of the steps matters.
// For example, if N is 4, then there are 5 unique ways:
//
// 1, 1, 1, 1
// 2, 1, 1
// 1, 2, 1
// 1, 1, 2
// 2, 2
//
// What if, instead of being able to climb 1 or 2 steps at a time, you could climb any number from a set of positive integers X? For example, if X = {1, 3, 5}, you could climb 1, 3, or 5 steps at a time.

// maxStepPerm will return the maximum number of ways the steps can be climed given the provided number of steps which can be skipped at each step.
// This function uses a greedy approach.
//
// runtime: O(n*len(steps)) because it iterates over 1..n then each step in steps
// space:   O(n) because of the dp array
func maxStepPerm(n int, steps []int) int {
	if n < 0 {
		return 0
	}

	dp := make([]int, n+1)
	dp[0] = 1

	for i := 1; i <= n; i++ {
		sum := 0
		for j := 0; j < len(steps); j++ {
			prev := i - steps[j]
			if steps[j] <= i && prev >= 0 {
				sum += dp[prev]
			}
		}

		dp[i] = sum
	}

	return dp[n]
}
