func lengthOfLongestSubstring(s string) int {
	charSet := make(map[byte]int)
	var left, result int

	for right := 0; right < len(s); right++ {
		// If it is a duplicate, then
		// we shrink left side of the window until
		// there's no duplicate anymore
		if i, ok := charSet[s[right]]; ok {
			// We update left border to the position after the duplicate.
			// Use max() in order to never move left border backwards, 
			// if there a duplicate in the very begining, e.g. "abba",
			// where once we encounter the second "a", it refers to i=0
			// as last encountered, which is less than the current left
			// border left=2.
			left = max(i + 1, left)
		}

		charSet[s[right]] = right

		length := right - left + 1
		if length > result {
			result = length
		}
	}

	return result
}

// Time:  O(n)
// Space: O(m), where m is the total number of unique characters in s