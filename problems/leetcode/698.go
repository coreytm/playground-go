package leetcode

import "sort"

// 698. Partition to K Equal Sum Subsets
// Given an array of integers nums and a positive integer k, find whether it's possible to divide this array into k non-empty subsets whose sums are all equal.
//
// Example 1:
// Input: nums = [4, 3, 2, 3, 5, 2, 1], k = 4
// Output: True
// Explanation: It's possible to divide it into 4 subsets (5), (1, 4), (2,3), (2,3) with equal sums.
//
// Note:
// 1 <= k <= len(nums) <= 16.
// 0 < nums[i] < 10000.

// kEqualSums will determine if the set of integers can be split into k subsets of equal sums.
// runtime: O(nlog(n))
// space: O(n)
func kEqualSums(in []int, k int) bool {
	// O(nlog(n))
	sort.Ints(in)

	totalSum := 0
	for _, num := range in {
		totalSum += num
	}

	// if totalSum doesn't divide k, then it's not possible
	if totalSum%k != 0 {
		return false
	}

	goalSum := totalSum / k
	out := []int{}
	sum := 0

	for i := len(in) - 1; i >= 0; i-- {
		sum += in[i]
		if sum > goalSum {
			sum -= in[i]
			out = append(out, in[i])
			continue
		}

		if sum == goalSum {
			out = append(out, in[0:i]...)
			break
		}
	}

	// none of the numbers add up to our goal sum
	if len(out) == len(in) {
		return false
	}

	// recurse on set of numbers after removing the summed set
	return kEqualSums(out, k-1)
}
