func productExceptSelf(nums []int) []int {
	prefix := make([]int, len(nums))
	suffix := make([]int, len(nums))
	prefix[0], suffix[len(nums)-1] = 1, 1

	for i := 1; i < len(nums); i++ {
		prefix[i] = nums[i-1] * prefix[i-1]
	}

	for i := len(nums) - 2; i >= 0; i-- {
		suffix[i] = nums[i+1] * suffix[i+1]
	}

	result := make([]int, 0, len(nums))
	for i := 0; i < len(nums); i++ {
		result = append(result, prefix[i] * suffix[i])
	}

	return result
}

// Time:  O(n)
// Space: O(n)
