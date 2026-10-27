func evalRPN(tokens []string) int {
	stack := make([]int, 0, len(tokens))

	for _, token := range tokens {
		if token == "+" {
			a := stack[len(stack)-2]
			b := stack[len(stack)-1]
			stack = stack[:len(stack)-2]

			stack = append(stack, a + b)
		} else if token == "-" {
			a := stack[len(stack)-2]
			b := stack[len(stack)-1]
			stack = stack[:len(stack)-2]

			stack = append(stack, a - b)
		} else if token == "*" {
			a := stack[len(stack)-2]
			b := stack[len(stack)-1]
			stack = stack[:len(stack)-2]

			stack = append(stack, a * b)
		} else if token == "/" {
			a := stack[len(stack)-2]
			b := stack[len(stack)-1]
			stack = stack[:len(stack)-2]

			stack = append(stack, a / b)
		} else {
			number, _ := strconv.Atoi(token)
			stack = append(stack, number)
		}
		
	}

	return stack[0]
}

// Time:  O(n)
// Space: O(n)