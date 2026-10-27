func isPalindrome(s string) bool {
	l, r := 0, len(s)-1

	for l < r {
		if !isAlphaNumeric(s[l]) {
			l++
			continue
		}

		if !isAlphaNumeric(s[r]) {
			r--
			continue
		}

		if toLower(s[l]) != toLower(s[r]) {
			return false
		}

		l++
		r--
	}

	return true
}

// Time:  O(n)
// Space: O(1)

func isAlphaNumeric(char byte) bool {
	// Is char within [a-z], [A-Z], and [0-9] ranges?
	return ('a' <= char && char <= 'z') ||
		('A' <= char && char <= 'Z') ||
		('0' <= char && char <= '9')
}

func toLower(char byte) byte {
	// Is char within [A-Z] range?
	if 'A' <= char && char <= 'Z' {
		return char + ('a' - 'A') // char + (97 - 65) == char + 32
	}

	return char
}