func isMonotonic(nums []int) bool {
    if len(nums) == 1 {
		return true
	}

	var isIncreasing, isSet bool

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			continue
		} else if nums[i] > nums[i-1] {
			if !isSet {
				isIncreasing = true
				isSet = true
			}

			if isSet && !isIncreasing {
				return false
			}
		} else if nums[i] < nums[i-1] {
			if !isSet {
				isIncreasing = false
				isSet = true
			}
			
			if isSet && isIncreasing {
				return false
			}
		}
	}

	return true
}

// Time:  O(n)
// Space: O(1)