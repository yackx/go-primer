// Our first program
package main

import "fmt"

func main() {
	//
	// Variable declaration
	//

	var a1 int // Declared. It has a default value of 0
	fmt.Printf("a1 declared: %d\n", a1)
	a1 = 10
	fmt.Printf("a1 assigned: %d\n", a1)

	a2 := 10             // Idiomatic
	fmt.Println(a2)      // We must use the value or code won't compile
	var a3 = 20          // var is not necessary and usually omitted
	fmt.Println(a3)      // We must use the value or code won't compile
	var b1, b2, b3 int   // Same type
	b1, b2, b3 = 1, 2, 3 // Assign multiple variables at once
	fmt.Println(b1, b2, b3)
	c1, c2, c3 := 1, 2, 3 // Multiple declarations
	fmt.Println(c1, c2, c3)
	const hello = "Hello, World!"
	fmt.Println(hello)
	const pi = 3.1415
	fmt.Println(pi)
}
