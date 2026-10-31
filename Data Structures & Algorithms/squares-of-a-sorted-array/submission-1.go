func sortedSquares(nums []int) []int {
	l, r, resIndex := 0, len(nums) - 1, len(nums) - 1

	result := make([]int, len(nums))

	for l <= r {
		lSqr := nums[l] * nums[l]
		rSqr := nums[r] * nums[r]

		if lSqr <= rSqr {
			result[resIndex] = rSqr
			r--
		} else {
			result[resIndex] = lSqr
			l++
		}

		resIndex--
	}

	return result
}

// Time:  O(n)
// Space: O(n)
