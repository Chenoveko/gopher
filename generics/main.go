package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

// ============================================================
// Use of any (Variadic Function)
// ============================================================
// Like with ...interface{}
func PrintList(list ...any) {
	for _, value := range list {
		fmt.Println(value)
	}
}

// ============================================================
// Types and Constraints
// ============================================================
type my_int int

func addition[T ~int | float64](numbers ...T) T {
	fmt.Printf("Numbers Type: %T | Numbers Value: %v\n", numbers, numbers)
	var total T
	for _, num := range numbers {
		total += num
	}
	return total
}

// ============================================================
// Create Constraints
// ============================================================
type restrictions interface {
	~int | ~float64 | ~float32
}

func addition_restricted[T restrictions](numbers ...T) T {
	fmt.Printf("Numbers Type: %T | Numbers Value: %v\n", numbers, numbers)
	var total T
	for _, num := range numbers {
		total += num
	}
	return total
}

// ============================================================
// Use Third Party Constraints
// ============================================================

func addition_third[T constraints.Integer | constraints.Float](numbers ...T) T {
	fmt.Printf("Numbers Type: %T | Numbers Value: %v\n", numbers, numbers)
	var total T
	for _, num := range numbers {
		total += num
	}
	return total
}

// ============================================================
// Constraints and Operators
// ============================================================
func Includes[T comparable](list []T, value T) bool {
	for _, item := range list {
		if item == value {
			return true
		}
	}
	return false
}

// ============================================================
// Generics Structs
// ============================================================
type Product[T uint | string] struct {
	Id T
	Desc string
	Price float32
}
func main() {
	// ============================================================
	// Invoke Functions with any
	// ============================================================
	PrintList("John", "Ken", 2, 4.3, true)

	// ============================================================
	// Invoke Functions with Types and Constraints
	// ============================================================
	fmt.Println(addition(12, 20, 40))
	fmt.Println(addition(12, 20.2, 40))
	var x my_int = 3
	fmt.Println(addition(x, 2, 40))
	fmt.Println(addition_restricted(x, 2, 40))
	fmt.Println(addition_third(5.3, 2, 40))

	// ============================================================
	// Invoke Functions with Constraints and Operators
	// ============================================================
	strings := []string{"a", "b", "c"}
	numbers := []int{1, 2, 3, 4, 5}

	fmt.Println(Includes(strings, "a"))
	fmt.Println(Includes(strings, "f"))
	fmt.Println(Includes(numbers, 5))

	// ============================================================
	// Instance Generic Struct
	// ============================================================
	product1 := Product[uint]{1, "Shoes", 50}
	product2 := Product[string]{"HDHHD-JDDHD", "Shoes", 50}

	fmt.Println(product1)
	fmt.Println(product2)

}
