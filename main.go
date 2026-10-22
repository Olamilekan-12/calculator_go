package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
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
	first, err := askNumber(scanner, "Enter the first number: ")
	if err != nil {
		fmt.Println()
		if !errors.Is(err, io.EOF) {
			fmt.Println("Error:", err)
		}
		return
	}
	op, err := askOperator(scanner, "Enter an operator (+, -, *, /): ")
	if err != nil {
		fmt.Println()
		if !errors.Is(err, io.EOF) {
			fmt.Println("Error:", err)
		}
		return
	}
	fmt.Printf("%g %s\n", first, op)
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

func parseOperator(text string) (Operator, bool) {
	op := Operator(strings.TrimSpace(text))
	switch op {
	case OpAdd, OpSubtract, OpMultiply, OpDivide:
		return op, true
	}
	return "", false
}

func askNumber(scanner *bufio.Scanner, prompt string) (float64, error) {
	for {
		fmt.Print(prompt)
		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				return 0, err
			}
			return 0, io.EOF
		}

		value, err := parseNumber(scanner.Text())
		if err != nil {
			fmt.Printf("%v. Try again.\n", err)
			continue
		}
		return value, nil
	}
}

func askOperator(scanner *bufio.Scanner, prompt string) (Operator, error) {
	for {
		fmt.Print(prompt)
		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				return "", err
			}
			return "", io.EOF
		}

		op, ok := parseOperator(scanner.Text())
		if !ok {
			fmt.Printf("%q is not an operator. Use + - * or /. Try again. \n", strings.TrimSpace(scanner.Text()))
			continue
		}
		return op, nil
	}
}
