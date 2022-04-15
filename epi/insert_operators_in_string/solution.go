package insert_operators_in_string

import "fmt"

type operator int

const (
	plus operator = iota
	multiply
)

func ExpressionSynthesis(digits []int, target int) bool {
	if len(digits) == 0 {
		return target == 0
	}
	return directedExpressionSynthesis(digits, target, 0, 0, []operator{}, []int{})
}

func directedExpressionSynthesis(digits []int, target int, pos int, currentValue int, operators []operator, operands []int) bool {
	currentValue = currentValue*10 + digits[pos]
	if pos == len(digits)-1 {
		// this is the last element in the list; evaluate the expression
		return evaluate(append(operands, currentValue), operators) == target
	}

	// try a plus operator at this position
	// we don't use the optimization from the book
	if directedExpressionSynthesis(digits, target, pos+1, 0, append(operators, plus), append(operands, currentValue)) {
		return true
	}

	// try a multiply operator at this position
	if directedExpressionSynthesis(digits, target, pos+1, 0, append(operators, multiply), append(operands, currentValue)) {
		return true
	}

	// try no operator at this position
	if directedExpressionSynthesis(digits, target, pos+1, currentValue, operators, operands) {
		return true
	}

	return false
}

func evaluate(operands []int, operators []operator) int {
	if len(operands) != len(operators)+1 {
		panic(fmt.Errorf("unexpected count of operators (%d) and operands (%d)", len(operands), len(operators)))
	}

	// multiplyEval will contain the values after all the multiply operators
	// are evaluated
	multiplyEval := make([]int, 0, len(operands))
	multiplyEval = append(multiplyEval, operands[0])

	for i := 0; i < len(operators); i++ {
		var last int
		if operators[i] == multiply {
			n := len(multiplyEval)
			multiplyEval, last = multiplyEval[:n-1], multiplyEval[n-1]
			multiplyEval = append(multiplyEval, last*operands[i+1])
		} else {
			multiplyEval = append(multiplyEval, operands[i+1])
		}
	}

	sumEval := 0
	for _, v := range multiplyEval {
		sumEval += v
	}

	return sumEval
}

func ExpressionSynthesis3PowN(digits []int, target int) bool {
	if len(digits) == 0 {
		return target == 0
	}
	return helper(digits, target, 1, digits[0], 0, digits[0], digits[0])
}

func helper(digits []int, target int, pos int, currentValue int, valUntilLastPlus int, valSinceLastPlus int, lastNumber int) bool {
	if pos >= len(digits) {
		return currentValue == target
	}

	// try plus
	{
		newValUntilLastPlus := currentValue
		newValSinceLastPlus := digits[pos]
		newLastNumber := digits[pos]
		newCurrentValue := newValUntilLastPlus + newValSinceLastPlus
		if helper(digits, target, pos+1, newCurrentValue, newValUntilLastPlus, newValSinceLastPlus, newLastNumber) {
			return true
		}
	}

	// try multiply
	{
		newValUntilLastPlus := valUntilLastPlus
		newValSinceLastPlus := valSinceLastPlus * digits[pos]
		newLastNumber := digits[pos]
		newCurrentValue := newValUntilLastPlus + newValSinceLastPlus
		if helper(digits, target, pos+1, newCurrentValue, newValUntilLastPlus, newValSinceLastPlus, newLastNumber) {
			return true
		}
	}

	// try concatenate
	{
		newValUntilLastPlus := valUntilLastPlus
		newLastNumber := lastNumber*10 + digits[pos]
		var newValSinceLastPlus int
		if lastNumber != 0 {
			newValSinceLastPlus = valSinceLastPlus / lastNumber * newLastNumber
		} else {
			newValSinceLastPlus = valSinceLastPlus * newLastNumber
		}
		newCurrentValue := newValUntilLastPlus + newValSinceLastPlus
		if helper(digits, target, pos+1, newCurrentValue, newValUntilLastPlus, newValSinceLastPlus, newLastNumber) {
			return true
		}
	}

	return false
}
