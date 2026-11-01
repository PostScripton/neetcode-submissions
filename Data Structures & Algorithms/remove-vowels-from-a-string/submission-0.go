func removeVowels(s string) string {
    var sb strings.Builder

	for _, char := range s {
		if isVowel(char) {
			continue
		}

		sb.WriteRune(char)
	}

	return sb.String()
}

func isVowel(char rune) bool {
	return char == 'a' ||
		char == 'e' ||
		char == 'i' ||
		char == 'o' ||
		char == 'u'
}

// Time:  O(n)
// Space: O(1)
