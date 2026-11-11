func lengthOfLastWord(s string) int {
	var result int
	var found bool

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if found {
				break
			}
			continue
		}

		found = true
		result++
	}

	return result
}

// Time:  O(n)
// Space: O(1)