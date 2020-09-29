package evaluate_rpn

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/stefantds/go-epi-judge/stack"
)

func Eval(expression string) int {
	expr := strings.Split(expression, ",")
	if len(expr) == 0 {
		return 0
	}

	evalStack := make(stack.Stack, 0)
	var operand1, operand2 interface{}

	for i := 0; i < len(expr); i++ {
		switch expr[i] {
		case "+", "-", "*", "/":
			evalStack, operand2 = evalStack.Pop()
			evalStack, operand1 = evalStack.Pop()
			operation := expr[i]
			evalStack = evalStack.Push(applyOp(operand1.(int), operation, operand2.(int)))
		default:
			valAsInt, err := strconv.Atoi(expr[i])
			if err != nil {
				panic(fmt.Errorf("expected an int: %s", err))
			}
			evalStack = evalStack.Push(valAsInt)
		}
	}

	_, result := evalStack.Pop()
	return result.(int)
}

func applyOp(a int, op string, b int) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		panic("unknown operation " + op)
	}
}
