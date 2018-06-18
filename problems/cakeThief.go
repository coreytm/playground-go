package problems

// This is a Knapsack Problem
//
// Cakes have a weight and a value.
// A Bag has a maximum weight capacity.
// Add cakes to the bag which will maximize the sum of their values.
//
// Assumptions:
//   - There may not be any cakes
//   - Bag weight capacity can be <= 0
//   - Sum of values or weights can be > math.MaxInt32

// Solution: Dynamic Programming

type (
	cake struct {
		weight int
		value  int
	}

	bag struct {
		capacity int
		value    int64
	}
)

func (b *bag) maxValue(cakes []cake) {
	if b.capacity <= 0 || len(cakes) == 0 {
		return
	}

	dp := make([]int64, b.capacity+1)
	dp[0] = 0

	for i := 1; i < b.capacity; i++ {
		for j := 0; j < len(cakes); j++ {
			if cakes[j].weight == 0 {
				return
			}

			if i-cakes[j].weight < 0 {
				continue
			}

			dp[i] = dp[i-1]

			cur := dp[i-cakes[j].weight] + int64(cakes[j].value)
			if cur > dp[i] {
				dp[i] = cur
			}
		}
	}

	b.value = dp[b.capacity]
}
