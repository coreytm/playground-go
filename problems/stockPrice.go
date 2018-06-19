package problems

import "math"

type (
	// stockPrices is an array of stock prices at time i
	stockPrices []int
)

func (in stockPrices) maxProfit() int {
	n := len(in)
	if n < 2 {
		return 0
	}

	dp := make([]int, n)
	dp[0] = 0

	min := math.MaxInt32

	for i := 1; i < len(in); i++ {
		if in[i-1] < min {
			min = in[i-1]
		}

		dp[i] = dp[i-1]
		if in[i]-min > dp[i] {
			dp[i] = in[i] - min
		}
	}

	return dp[n]
}
