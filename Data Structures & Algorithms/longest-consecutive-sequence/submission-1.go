func longestConsecutive(nums []int) int {
	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}

	var result int

	for num := range numSet {
		length := 1

		// Is it a start of a sequence?
		if !numSet[num - 1] {
			for {
				if numSet[num + length] {
					length++
				} else {
					break
				}
			}

			if length > result {
				result = length
			}
		}
	}

	return result
}

// Time:  O(n)
// Space: O(n)