package main

import (
	"fmt"
)

func main() {
	// An array of integers, fixed-size. All values default to 0
	var x [5]int
	// Modify the first element as position 0
	x[0] = 1
	fmt.Println(x)

	// Initiliaze an array
	y := [5]int{1, 2, 3, 4, 5}
	fmt.Println(y)

	// If you use multiple lines,
	// the comma is mandatory after the last element
	z := [3]int{
		1,
		2,
		3,
	}
	fmt.Println(z)

	// A slice of integers.
	// Slices are idiomatic in Go (rather than using arrays)
	var sl []int
	// It's empty
	fmt.Println(sl)
	// Let's fill it with some values
	for i := 0; i <= 5; i++ {
		sl = append(sl, i)
	}
	fmt.Println("sl", sl)

	// Access an element
	fmt.Println(sl[3])

	// Slice it!
	fmt.Println(sl[0:2]) // [0, 1]
	fmt.Println(sl[:3])  // [0, 1, 2]
	fmt.Println(sl[3:])  // [3, 4, 5]

	// Delete = slice and append.
	// Notice the use of the variadic operator "..." as a second argument
	sl = append(sl[:3], sl[4:]...)
	fmt.Println("sl after cut", sl)

	// Copy
	clone := make([]int, len(sl))
	copy(clone, sl)
	fmt.Println("clone", clone)
}
