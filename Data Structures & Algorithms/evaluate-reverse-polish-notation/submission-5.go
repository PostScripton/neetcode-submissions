var operatorsMap = map[string]struct{}{
	"+": struct{}{},
	"-": struct{}{},
	"*": struct{}{},
	"/": struct{}{},
}

func evalRPN(tokens []string) int {
	operands := make([]int, 0, 2)
	var operator string

	for _, token := range tokens {
		if _, ok := operatorsMap[token]; ok {
			operator = token
		} else {
			number, _ := strconv.Atoi(token)
			operands = append(operands, number)
		}

		if len(operands) >= 2 && operator != "" {
			result := eval(operands[len(operands)-2], operands[len(operands)-1], operator)

			operands = operands[:len(operands)-2]
			operands = append(operands, result)

			operator = ""
		}
	}

	return operands[0]
}

func eval(a, b int, operator string) int {
	switch operator {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		return 0
	}
}