package main

import (
	"errors"
	"fmt"
)

func main() {

	a := 10.0
	b := 4.0
	fmt.Printf("%g + %g = %g\n", a, b, add(a, b))
	fmt.Printf("%g - %g = %g\n", a, b, subtract(a, b))
	fmt.Printf("%g * %g = %g\n", a, b, multiply(a, b))
	result, err := divide(a, b)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Printf("%g / %g = %g\n", a, b, result)
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
