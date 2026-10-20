package main

import "fmt"

func main() {

	a := 10.0
	b := 4.0
	fmt.Printf("%g + %g = %g\n", a, b, add(a, b))
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

func divide(a, b float64) float64 {
	return a / b
}
