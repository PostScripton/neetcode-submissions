func sortedSquares(nums []int) []int {
	l, r := 0, len(nums) - 1

	result := make([]int, 0, len(nums))

	for l <= r {
		lSqr := nums[l] * nums[l]
		rSqr := nums[r] * nums[r]

		if lSqr <= rSqr {
			result = append(result, rSqr)
			r--
		} else {
			result = append(result, lSqr)
			l++
		}
	}

	for i := 0; i < len(result) / 2; i++ {
		left, right := i, len(result) - 1 - i
		result[left], result[right] = result[right], result[left]
	}

	return result
}

// Time:  O(n)
// Space: O(n)
