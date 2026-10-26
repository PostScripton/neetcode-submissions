func search(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	mid := -1

	for {
		if nums[l] == target {
			return l
		}
		if nums[r] == target {
			return r
		}

		mid = (l + r) / 2

		if nums[mid] == target {
			return mid
		}

		// Comparing the last 2 elements no longer makes sense
		if mid == l {
			break
		}
		
		if nums[mid] < target {
			l = mid
		} else if nums[mid] > target {
			r = mid
		}
	}

	return -1
}
