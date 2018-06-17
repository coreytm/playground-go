package problems

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

func (nums nums) findDuplicate() int {
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
