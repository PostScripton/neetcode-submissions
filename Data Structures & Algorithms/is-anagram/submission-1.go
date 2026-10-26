func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sMap := make(map[rune]int)

	// Init map for s string
	for _, char := range s {
		sMap[char]++
	}

	// Check if all characters from t exist in the map
	for _, char := range t {
		if freq, ok := sMap[char]; ok && freq > 0 {
			sMap[char]--
			continue
		}

		return false
	}

	return true
}
