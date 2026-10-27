var operatorsMap = map[string]struct{}{
	"+": struct{}{},
	"-": struct{}{},
	"*": struct{}{},
	"/": struct{}{},
}

func evalRPN(tokens []string) int {
	operands := make([]int, 0, 2)
	operators := make([]string, 0, 1)

	for _, token := range tokens {
		if _, ok := operatorsMap[token]; ok {
			operators = append(operators, token)
		} else {
			number, _ := strconv.Atoi(token)
			operands = append(operands, number)
		}

		if len(operands) >= 2 && len(operators) >= 1 {
			result := eval(operands[len(operands)-2], operands[len(operands)-1], operators[len(operators)-1])

			operands = operands[:len(operands)-2]
			operands = append(operands, result)

			operators = operators[:len(operators)-1]
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