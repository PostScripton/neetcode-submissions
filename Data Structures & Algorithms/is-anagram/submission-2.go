func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	charsMap := make(map[rune]int)

	// Init map with all used characters
	for i := range s {
		charsMap[rune(s[i])]++
		charsMap[rune(t[i])]--
	}

	// Check if any character is not self-liquidated
	for _, freq := range charsMap {
		if freq != 0 {
			return false
		}
	}

	return true
}
