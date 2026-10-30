func isConsecutive(nums []int) bool {
	numSet := make(map[int]bool)
	minNum := nums[0]
	maxNum := nums[0]

	for _, num := range nums {
		if numSet[num] {
			return false
		}

		if num < minNum {
			minNum = num
		}
		if num > maxNum {
			maxNum = num
		}

		numSet[num] = true
	}

	return maxNum - minNum + 1 == len(nums)
}

// Time:  O(n)
// Space: O(n)
