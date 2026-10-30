func isConsecutive(nums []int) bool {
	numSet := make(map[int]bool)
	minNum := nums[0]

	for _, num := range nums {
		if numSet[num] {
			return false
		}

		if num < minNum {
			minNum = num
		}
		numSet[num] = true
	}

	for num := minNum; num < minNum + len(nums); num++ {
		if !numSet[num] {
			return false
		}
	}

	return true
}

// Time:  O(n)
// Space: O(n)
