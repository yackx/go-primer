package main

import "fmt"

func main() {
	// Iterate over a slice with range
	numbers := []int{6, 4, 2}
	for index, value := range numbers {
		fmt.Printf("%d => %d\n", index, value)
	}

	// The classic for loop.
	// Range is your first choice but do not hesitate to opt for this proven construct
	for i := 0; i < len(numbers); i++ {
		fmt.Printf("%d => %d\n", i, numbers[i])
	}

	// Sum numbers. The index is ignored
	var sum int
	for _, value := range numbers {
		sum += value
	}
	fmt.Println("sum", sum)

	// Iterate over a map
	fruits := map[string]int{"banana": 2, "apple": 5}
	fmt.Println("fruits", fruits)
	for k, v := range fruits {
		fmt.Printf("k=%s; v=%s\n", k, v)
	}

	// Map keys only
	for k := range fruits {
		fmt.Printf("k=%s\n", k)
	}

	// range on strings iterates over Unicode code points.
	// The first value is the index
	// and the second the "rune"
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
