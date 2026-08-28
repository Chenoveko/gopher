package main

import "fmt"

func edit_number(pointer *int) {
	*pointer = 20
}

// ============================================================
// Define Structs
// ============================================================
type Person struct {
	name     string
	age      int
	email    string
	is_adult bool
}

// ============================================================
// Define Methods
// ============================================================
func (person *Person) sayHello() {
	fmt.Println("Hi, my name is ", person.name)
}

// ============================================================
// Custom Types
// ============================================================

func main() {
	// ============================================================
	// Instance Struct
	// ============================================================
	var go_creator1 Person

	go_creator1.name = "Ken"
	go_creator1.age = 83
	go_creator1.email = "kthopmson@gmail.com"
	go_creator1.is_adult = true

	go_creator2 := Person{"Robert", 78, "rgriesemer@gmail.com", true}
	go_creator2.age = 80

	fmt.Println("Go Creator 1 Struct: ", go_creator1)
	fmt.Println("Go Creator 2 Struct: ", go_creator2)

	// ============================================================
	// Pointers
	// ============================================================
	var number int = 10
	var pointer_to_number *int = &number

	fmt.Println("Number: ", number)
	fmt.Println("Pointer to number: ", pointer_to_number)

	// We pass the pointet to edit the value. If we pass the var, we will send a copy to the function
	edit_number(pointer_to_number)
	fmt.Println("Edit number: ", number)

	// ============================================================
	// Invoke Methods
	// ============================================================
	go_creator1.sayHello()
}
