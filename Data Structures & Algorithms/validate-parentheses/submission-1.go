var parentheses = map[rune]rune{
	'(': ')',
	'{': '}',
	'[': ']',
}

func isValid(s string) bool {
	// Check that amount of symbols is even,
	// corresponding to open and closed parentheses
	if len(s) % 2 != 0 {
		return false
	}

	// A stack of closing parentheses
    stack := make([]rune, 0, len(s))

	for _, char := range s {
		if closing, open := parentheses[char]; open {
			stack = append(stack, closing)
			continue
		}

		if len(stack) > 0 && stack[len(stack)-1] == char {
			// The top stack element is equal to the current closing parenthesis
			// It is valid -> remove from stack
			stack = stack[:len(stack)-1]
		} else {
			// Invalid sequence
			return false
		}
	}

	return len(stack) == 0
}
