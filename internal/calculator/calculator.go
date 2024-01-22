package calculator

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/Avik32223/coding-challenges-solutions/pkg/slices"
)

func getPrecedence(c rune) int {
	if c == '^' {
		return 3
	} else if c == '/' || c == '*' {
		return 2
	} else if c == '+' || c == '-' {
		return 1
	}
	return -1
}

func getAssociativity(c rune) rune {
	if c == '^' {
		return 'R'
	}
	return 'L'
}

func infixToPostfix(source string) []string {
	result := make([]string, 0)
	operatorStack := make(slices.Stack[rune], 0)
	for _, i := range strings.Split(source, " ") {
		if _, err := strconv.ParseFloat(i, 64); err != nil {
			for _, op := range i {
				if op == '(' {
					operatorStack = append(operatorStack, op)
				} else if op == ')' {
					for len(operatorStack) > 0 {
						opLast := operatorStack[len(operatorStack)-1]
						if opLast != '(' {
							result = append(result, string(operatorStack.Pop()))
						} else {
							operatorStack.Pop()
							break
						}
					}
				} else {
					for len(operatorStack) > 0 {
						opLast := operatorStack[len(operatorStack)-1]
						if (getPrecedence(op) < getPrecedence(opLast)) || (getPrecedence(op) == getPrecedence(opLast) && getAssociativity(op) == 'L') {
							result = append(result, string(operatorStack.Pop()))
						} else {
							break
						}
					}
					operatorStack = append(operatorStack, op)
				}
			}
		} else {
			result = append(result, i)
		}
	}

	for len(operatorStack) > 0 {
		result = append(result, string(operatorStack.Pop()))
	}

	return result
}

func evaluatePostfix(postfix []string) (result string) {
	valStack := slices.Stack[float64]{}
	for _, i := range postfix {
		if val, err := strconv.ParseFloat(i, 64); err != nil {
			val1 := valStack.Pop()
			val2 := valStack.Pop()
			var result float64
			switch i {
			case "^":
				result = math.Pow(val2, val1)
			case "*":
				result = val2 * val1
			case "/":
				result = val2 / val1
			case "+":
				result = val2 + val1
			case "-":
				result = val2 - val1
			default:
				panic("Unknown operand " + i)
			}
			valStack = append(valStack, result)
		} else {
			valStack = append(valStack, val)
		}
	}
	val := valStack.Pop()
	if val == float64(int(val)) {
		result = fmt.Sprintf("%d", int(val))
	} else {
		result = fmt.Sprintf("%.2f", val)
	}
	return
}

// Calculate works with space separated strings.
// It evaluates mathematical expressions with binary operators (+,-,/,*,^)
func Calculate(source string) (string, error) {
	postfixNotation := infixToPostfix(source)
	result := evaluatePostfix(postfixNotation)
	return result, nil
}
