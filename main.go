package main

import "fmt"

func main() {
	a := float64(10)
	b := 4.0

	fmt.Printf("%g + %g = %g\n", a, b, a+b)
	fmt.Printf("%g - %g = %g\n", a, b, a-b)
	fmt.Printf("%g * %g = %g\n", a, b, a*b)
	fmt.Printf("%g / %g = %g\n", a, b, a/b)
}
