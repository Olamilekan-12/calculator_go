package operations

import (
	"errors"
	"fmt"
	"strings"
)

type Operator string

const (
	OpAdd      Operator = "+"
	OpSubtract Operator = "-"
	OpMultiply Operator = "*"
	OpDivide   Operator = "/"
)

func ParseOperator(text string) (Operator, bool) {
	op := Operator(strings.TrimSpace(text))
	switch op {
	case OpAdd, OpSubtract, OpMultiply, OpDivide:
		return op, true
	}
	return "", false
}

func Calculate(a float64, op Operator, b float64) (float64, error) {
	switch op {
	case OpAdd:
		return add(a, b), nil
	case OpSubtract:
		return subtract(a, b), nil
	case OpMultiply:
		return multiply(a, b), nil
	case OpDivide:
		return divide(a, b)
	default:
		return 0, fmt.Errorf("unknown operator %q", op)
	}
}

func add(a, b float64) float64 {
	return a + b
}

func subtract(a, b float64) float64 {
	return a - b
}

func multiply(a, b float64) float64 {
	return a * b
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}
