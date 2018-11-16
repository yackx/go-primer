package main

import "fmt"

func main() {
	// Map (k=string, v=int)
	m := make(map[string]int)
	m["foo"] = 2
	m["bar"] = 3
	fmt.Println("map:", m)

	// Initialize a new map
	fruits := map[string]int{"banana": 2, "apple": 5}
	fmt.Println("fruits", fruits)

	// How many elements in the map
	fmt.Println("len(fruits)", len(fruits))

	// Retrieve element
	orange := m["orange"]
	fmt.Println("orange", orange)

	// Element not found: default value returned (0 for int)
	cucumber := m["cucumber"]
	fmt.Println("cucumber", cucumber)

	// Test if element is present. Notice the idiomatic construct
	blueberry, ok := m["blueberry"]
	if ok {
		fmt.Println("Found blueberry", blueberry) // not going here
	} else {
		fmt.Println("blueberry not found") // going here
	}

	// You can drop the first variable if you don't need it. Use '_'
	_, ok = m["blueberry"]
	if !ok {
		fmt.Println("blueberry not found")
	}

	// Delete
	delete(fruits, "banana")
	_, ok = m["banana"]
	if !ok {
		fmt.Println("banana not found")
	}
}
