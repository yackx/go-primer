package main

import (
	"errors"
	"fmt"
	"os"
)

// A simple function
func hello() {
	fmt.Println("Hello")
}

// A function with 1 parameter
func sayHello(to string) {
	fmt.Println("Hello, ", to)
}

// A function with 2 parameters that returns one value
func sum(a, b int) int {
	return a + b
}

// Return a value and an error.
// Either of them can be 'nil'
func div(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by 0")
	}
	return a / b, nil
}

// Variadic parameter '...'
// Add a variable number of integers.
// Notice the return value of int with a variable "sum" declared upfront
func add(x ...int) (sum int) {
	for _, n := range x {
		sum += n
	}
	return // implicitly: 'sum' declared as a return value
}

func main() {
	hello()
	sayHello("Go")

	fmt.Println(sum(3, 4))

	x := add(1, 2, 3)
	fmt.Println("x", x)

	d, err := div(10, 5)
	if err != nil {
		panic("Failed to divide") // should not happen
	}
	fmt.Println("10 / 5 = ", d)

	_, err = div(6, 0)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	}
}
