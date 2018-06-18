package problems

// Find a duplicate, Space Edition.
//
// We have an array of integers, where:
//   - The integers are in the range 1..n
//   - The array has a length of n+1
//   - It follows that our array has at least one integer which appears at least twice. But it may have several duplicates, and each duplicate may appear more than twice.
//
// Write a method which finds an integer that appears more than once in our array. (If there are multiple duplicates, you only need to find one of them.)
//
// ----------
//
// Note: This solution only works if the following assumptions are met.
//
// Assumptions:
//   - len(nums) = n - 1
//   - nums[i] is in 1..n
//
// ex. n = 8, len(nums) = 9, nums[i] is in 1..8
//
// runtime: O(nlog(n))
// space:   O(1)

type (
	nums []int
)

// runtime: O(n)
// space  : O(n)
func (nums nums) findDuplicateUsingHashMap() int {
	hm := map[int]bool{}

	for i := 0; i < len(nums); i++ {
		if hm[nums[i]] {
			return nums[i]
		}

		hm[nums[i]] = true
	}

	return -1
}

// runtime: O(nlog(n))
// space:   O(1)
func (nums nums) findDuplicateOptimizedForSpace() int {
	floor := 1
	ceiling := len(nums) - 1

	for floor < ceiling {
		middle := floor + int(((ceiling - floor) / 2))
		lowerFloor := floor
		lowerCeiling := middle
		upperFloor := middle + 1
		upperCeiling := ceiling

		cntInLowerRange := 0
		for _, num := range nums {
			if num >= lowerFloor && num <= lowerCeiling {
				cntInLowerRange++
			}
		}

		cntOfDistinctInLowerRange := lowerCeiling - lowerFloor + 1
		if cntInLowerRange > cntOfDistinctInLowerRange {
			floor = lowerFloor
			ceiling = lowerCeiling
		} else {
			floor = upperFloor
			ceiling = upperCeiling
		}
	}

	return floor
}

// findDuplicateOptimizedForSpaceBeastMode optimizes both runtime and space
// runtime: O(n)
// space:   O(1)
//
// 1) find cycle
// 2) find length of cycle
// 3) two pointers spaced out by length of cycle starting at n
// 4) shift both over until they converge
func (nums nums) findDuplicateOptimizedForSpaceBeastMode() int {
	n := len(nums)

	p1 := n + 1
	for i := 0; i < n; i++ {
		p1 = nums[p1-1]
	}

	p2 := nums[p1-1]
	cnt := 1
	for p1 != p2 {
		p2 = nums[p2-1]
		cnt++
	}

	p2 = n + 1
	for i := 0; i < cnt; i++ {
		p2 = nums[p2-1]
	}

	p1 = n + 1
	for p1 != p2 {
		p1 = nums[p1-1]
		p2 = nums[p2-2]
	}

	return p1
}
