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

	"calculator/operations"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	if err := run(scanner); err != nil {
		fmt.Println()
		if !errors.Is(err, io.EOF) {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
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

func askOperator(scanner *bufio.Scanner, prompt string) (operations.Operator, error) {
	for {
		fmt.Print(prompt)
		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				return "", err
			}
			return "", io.EOF
		}

		op, ok := operations.ParseOperator(scanner.Text())
		if !ok {
			fmt.Printf("%q is not an operator. Use + - * or /. Try again. \n", strings.TrimSpace(scanner.Text()))
			continue
		}
		return op, nil
	}
}

func runCalculation(scanner *bufio.Scanner) error {
	first, err := askNumber(scanner, "Enter the first number: ")
	if err != nil {
		return err
	}

	op, err := askOperator(scanner, "Enter an operator (+, -, *, /): ")
	if err != nil {
		return err
	}

	second, err := askNumber(scanner, "Enter the second number: ")
	if err != nil {
		return err
	}

	result, err := operations.Calculate(first, op, second)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	fmt.Printf("%g %s %g = %g\n", first, op, second, result)
	return nil
}

func askAgain(scanner *bufio.Scanner) (bool, error) {
	fmt.Print("Calculate again? (y/n): ")
	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return false, err
		}
		return false, io.EOF
	}
	answer := strings.ToLower(strings.TrimSpace(scanner.Text()))
	return answer == "y", nil
}

func run(scanner *bufio.Scanner) error {
	for {
		if err := runCalculation(scanner); err != nil {
			return err
		}

		again, err := askAgain(scanner)
		if err != nil {
			return err
		}
		if !again {
			return nil
		}
	}
}
