package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Operator string

const (
	OpAdd      Operator = "+"
	OpSubtract Operator = "-"
	OpMultiply Operator = "*"
	OpDivide   Operator = "/"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter the first number: ")
	scanner.Scan()
	answer := scanner.Text()
	first, err := parseNumber(answer)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Printf("%g (%T)\n", first, first)
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

func parseNumber(text string) (float64, error) {
	trimmed := strings.TrimSpace(text)
	if trimmed == "" {
		return 0, errors.New("please enter a number")
	}

	value, err := strconv.ParseFloat(trimmed, 64)
	if err != nil || math.IsNaN(value) || math.IsInf(value, 0) {
		return 0, fmt.Errorf("%q is not a valid number", trimmed)
	}

	return value, nil
}
