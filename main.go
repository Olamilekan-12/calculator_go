package main

import (
	"errors"
	"fmt"
)

type Operator string

const (
	OpAdd      Operator = "+"
	OpSubtract Operator = "-"
	OpMultiply Operator = "*"
	OpDivide   Operator = "/"
)

func main() {
	result, err := calculate(10, OpDivide, 4)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println(result)
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

func calculate(a float64, op Operator, b float64) (float64, error) {
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
