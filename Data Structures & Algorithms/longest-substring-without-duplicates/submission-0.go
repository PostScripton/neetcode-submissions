func lengthOfLongestSubstring(s string) int {
	charSet := make(map[byte]bool)
	var left, max int

	for right := 0; right < len(s); right++ {
		// If it is a duplicate, then
		// we shrink left side of the window until
		// there's no duplicate anymore
		for charSet[s[right]] {
			delete(charSet, s[left])
			left++
		}

		charSet[s[right]] = true

		length := right - left + 1
		if length > max {
			max = length
		}
	}

	return max
}

// Time:  O(n)
// Space: O(m), where m is the total number of unique characters in s