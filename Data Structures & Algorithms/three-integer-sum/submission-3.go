func threeSum(nums []int) [][]int {
	sort.Ints(nums) // Time: O(n log n); Space: O(1)

	result := make([][]int, 0, len(nums) / 3)

	for i := 0; i < len(nums)-2; i++ {
		// Skip duplicates because the input is sorted
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i + 1, len(nums) - 1

		for left < right {
			threeSum := nums[i] + nums[left] + nums[right]
			
			if threeSum < 0 {
				left++
			} else if threeSum > 0 {
				right--
			} else {
				result = append(result, []int{nums[i], nums[left], nums[right]})

				left++
				right--

				// Skip duplicates because the input is sorted
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			}
		}
	}

	return result
}

// Time:  O(n^2)
// Space: O(1)
