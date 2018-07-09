package daily

// This problem was asked by Facebook.
// A builder is looking to build a row of N houses that can be of K different colors.
// He has a goal of minimizing cost while ensuring that no two neighboring houses are of the same color.
//
// Given an N by K matrix where the nth row and kth column represents the cost to build the nth house with kth color, return the minimum cost which achieves this goal.

// minCost will return the minimun cost to paint all houses where no adjacent houses are the same colour.
// Dynamic programming approach: Keep track of the cost to paint each house each colour.
//
// runtime: O(nk^2) because O(nk) just to itereate through matrix, then O(k) for each iteration of k when finding min.
// space: O(1) or O(n) if you count the array used for getting min.
func minCost(costs [][]int) int {
	n := len(costs)
	if n == 0 {
		return 0
	}
	k := len(costs[0])
	if k == 0 {
		return 0
	}

	for i := 1; i < len(costs); i++ {
		for j := 0; j < k; j++ {
			costs[i][j] += min(append(costs[i-1][:j], costs[i-1][j+1:]...))
		}
	}

	return min(costs[n-1])
}

func min(in []int) int {
	min := 0
	for _, num := range in {
		if num < min {
			min = num
		}
	}

	return min
}
