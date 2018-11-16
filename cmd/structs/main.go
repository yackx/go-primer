package main

import "fmt"

type Car struct {
	Brand string
	Model string
	Doors int
}

func main() {
	audi := Car{"Audi", "A4", 5}
	fmt.Println(audi)

	bmw := Car{Brand: "BMW", Model: "325i", Doors: 4}
	fmt.Println(bmw)
	fmt.Println("Brand:", bmw.Brand)
}
