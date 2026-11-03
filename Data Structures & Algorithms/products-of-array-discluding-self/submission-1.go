func productExceptSelf(nums []int) []int {
	prefix, suffix := 1, 1
	result := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		result[i] = prefix
		prefix *= nums[i]
	}

	for i := len(nums) - 1; i >= 0; i-- {
		result[i] *= suffix
		suffix *= nums[i]
	}

	return result
}

// Time:  O(n)
// Space: O(1)
