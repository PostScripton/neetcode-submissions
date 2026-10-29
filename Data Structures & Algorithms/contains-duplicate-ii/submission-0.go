func containsNearbyDuplicate(nums []int, k int) bool {
	if k == 0 {
		return false
	}

	set := make(map[int]bool)

	l, r := 0, 0

	for r < len(nums) {
		length := r - l
		if length > k {
			delete(set, nums[l])
			l++
			continue
		}

		if set[nums[r]] {
			return true
		}

		set[nums[r]] = true
		r++
	}

	return false
}

// Time:  O(n)
// Space: O(k)
