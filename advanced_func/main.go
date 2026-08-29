package main

import "fmt"

// ============================================================
// Variadic Functions
// ============================================================

func addition(numbers ...int) int {
	fmt.Printf("Numbers Type: %T | Numbers Value: %v\n",numbers, numbers)
	var total int 
	for _, num := range numbers {
		total += num
	}
	return total
}

func displayData(data ...interface{}) {
	for _, d := range data {
		fmt.Printf("Data Type: %T | Data Value: %v\n",d, d)
	}
}

// ============================================================
// Recursive Functions
// ============================================================
func factorial(number int) int {
	if number == 0 {
		return 1
	}
	return number * factorial(number-1)
}

// ============================================================
// Higher-Order Functions
// ============================================================
// Passing functions as arguments

func greetings(name string, f func(string)) {
	f(name)
}

func double(f func(int) int, x int) int {
	return f(x * 2)
}

func addOne(x int) int {
	return x + 1
}

// ============================================================
// Closures
// ============================================================
func incrementer() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	// ============================================================
	// Invoke Variadic Functions
	// ============================================================
	fmt.Println(addition(12, 20, 40))
	displayData("Hello", true, 45)

	// ============================================================
	// Invoke Recursive Functions
	// ============================================================
	fmt.Println(factorial(3))

	// ============================================================
	// Anonymous Functions
	// ============================================================
	func () {
		fmt.Println("I am an anonymous function")
	}() // Execute

	// ============================================================
	// Invoke Higher-Order Functions
	// ============================================================
	greet := func(name string) {
		fmt.Printf("Hi, %s\n", name)
	}
	greetings("John", greet)

	r := double(addOne, 3)
	fmt.Println("Result: ", r)

	// ============================================================
	// Invoke Closures
	// ============================================================
	nextInt := incrementer()
	fmt.Println("Closure: ", nextInt())
	fmt.Println("Closure: ", nextInt())

}
