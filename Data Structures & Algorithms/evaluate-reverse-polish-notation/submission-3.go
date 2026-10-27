type Operator string

const (
	OperatorUnknown Operator = "?"
	OperatorAdd Operator = "+"
	OperatorSub Operator = "-"
	OperatorMul Operator = "*"
	OperatorDiv Operator = "/"
)

func evalRPN(tokens []string) int {
	operands := make([]int, 0, 2)
	operators := make([]Operator, 0, 1)

	for _, token := range tokens {
		if isOperator(token) {
			operators = append(operators, Operator(token))
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

func isOperator(char string) bool {
	switch Operator(char) {
	case OperatorAdd:
		fallthrough
	case OperatorSub:
		fallthrough
	case OperatorMul:
		fallthrough
	case OperatorDiv:
		return true
	default:
		return false
	}
}

func eval(a, b int, operator Operator) int {
	switch operator {
	case OperatorAdd:
		return a + b
	case OperatorSub:
		return a - b
	case OperatorMul:
		return a * b
	case OperatorDiv:
		return a / b
	default:
		return 0
	}
}