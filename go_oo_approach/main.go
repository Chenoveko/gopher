package main

import (
	"go_oo_approach/pkg/animal"
	"go_oo_approach/pkg/library"
)

func main() {
	quijote := library.NewBook("Quijote", "Cervantes", 986)
	quijote.PrintInfo()
	workbook := library.NewTextBook("Cambridge", "John", 45, "Oxford", "C1")
	workbook.PrintInfo()

	// Using interfaces
	library.Print(quijote)

	myDog := animal.Dog{Name: "Scooby"}
	myCat := animal.Cat{Name: "Garfield"}

	animal.MakeSound(&myDog)
	animal.MakeSound(&myCat)

}
