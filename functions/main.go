package main

import (
	"fmt"
	"math/rand"
)

// ============================================================
// Define Functions
// ============================================================
// Function with parameters and no return value
func hello(name string) {
	fmt.Println("Hello, ", name)
}

// Function that returns a single value
func even_or_odd(number int) bool {
	if number%2 == 0 {
		return true
	} else {
		return false
	}
}

// Function that returns multiple values
func calculation(x, y int) (int, int) {
	addition := x + y
	multiplication := x * y
	return addition, multiplication
}

// Function with named return values
func same_calculation(x, y int) (addition, multiplication int) {
	addition = x + y
	multiplication = x * y
	return
}

// Return a random integer between min and max
func random_number(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func main() {
	// ============================================================
	// Invoke Functions
	// ============================================================
	hello("John")
	fmt.Println("Is 2 even?", even_or_odd(2))
	fmt.Println("Is 3 even?", even_or_odd(3))

	addition, multiplication := calculation(2, 3)

	fmt.Println("Addition:", addition)
	fmt.Println("Multiplication:", multiplication)

	addition, multiplication = same_calculation(4, 5)

	fmt.Println("Named return addition:", addition)
	fmt.Println("Named return multiplication:", multiplication)

	// Generate a random number between 1 and 100
	number := random_number(1, 100)
	fmt.Println("Random number:", number)

}
