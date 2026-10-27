func threeSum(nums []int) [][]int {
	sort.Ints(nums) // Time: O(n log n)

	triplets := make(map[[3]int]struct{})

	i, j, k := 0, 1, len(nums)-1
	for i < len(nums)-2 {
		// We transform equasion:
		// 1. nums[i] + nums[j] + nums[k] == 0
		// 2. nums[i] == -(nums[j] + nums[k])
		// 3. -nums[i] == nums[j] + nums[k]
		//
		// And then with 2 right pointers (j & k) search for Two Sum to get target i

		if j >= k {
			i++
			j = i + 1
			k = len(nums) - 1
		}

		if nums[j] + nums[k] < -nums[i] {
			j++
		} else if nums[j] + nums[k] > -nums[i] {
			k--
		} else {
			triplets[[3]int{nums[i], nums[j], nums[k]}] = struct{}{}
			j++
			k--
		}
	}

	result := make([][]int, 0, len(triplets))
	for triplet := range triplets {
		result = append(result, triplet[:])
	}

	return result
}

// Time:  O(n^2)
// Space: O(n)
