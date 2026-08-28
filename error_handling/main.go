package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

// ============================================================
// Define Errors
// ============================================================
// Create an error with a fixed message
func divide(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return x / y, nil
}

// Create a formatted error message
func divide_fmt(x, y int) (int, error) {
	if y == 0 {
		return 0, fmt.Errorf("cannot divide by %d", y)
	}
	return x / y, nil
}

// ============================================================
// Panic Handling with Recover
// ============================================================
// recover catches a panic and prevents the program from crashing
func divide_panic(x, y int) {
	// Anonymous function executed when divide_panic returns
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	validate_zero(y)
	fmt.Println("Divide Panic Result:", x/y)
}

func validate_zero(x int) {
	if x == 0 {
		panic("cannot divide by zero")
	}
}

func main() {
	// ============================================================
	// Error Handling
	// ============================================================
	// We use errors for expected errors
	my_string := "123"
	num, err := strconv.Atoi(my_string)
	if err != nil {
		fmt.Println("Error: ", err)
		return // Exit main
	}
	fmt.Println("Number: ", num)

	division, division_error := divide(10, 2)
	if division_error != nil {
		fmt.Println("Division Error: ", division_error)
		return // Exit main
	}
	fmt.Println("Division: ", division)

	division, division_error = divide_fmt(10, 1)
	if division_error != nil {
		fmt.Println("Division Error fmt: ", division_error)
		return // Exit main
	}
	fmt.Println("Division fmt: ", division)

	// ============================================================
	// Defer
	// ============================================================
	// Deferred functions are executed when the current function returns
	// Multiple defer statements follow LIFO order (Last In, First Out)
	// Use defer to clean up resources such as files
	defer fmt.Println(3)
	defer fmt.Println(1)
	fmt.Println(2)

	// Example with file
	file, file_err := os.Create("file.txt")
	if file_err != nil {
		fmt.Println("os.Create Error: ", file_err)
		return // Exit main
	}
	defer file.Close() // Close the file when main returns
	_, file_err = file.Write([]byte("Hello World!"))
	if file_err != nil {
		fmt.Println("file.Write Error: ", file_err)
		return // Exit main
	}

	// ============================================================
	// Panic and Recover
	// ============================================================
	// Use errors for expected problems
	// Use panic only for serious or unrecoverable situations.
	fmt.Println("Panic division examples:")
	divide_panic(2, 1)
	divide_panic(2, 9)
	divide_panic(2, 0)

	// ============================================================
	// Loggin Errors
	// ============================================================
	log.SetPrefix("main(): ")
	log.Println("Log")

	// Save logs in file
	log_file, log_err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if log_err != nil {
		log.Fatal(log_err)
	}
	defer log_file.Close()
	log.SetOutput(log_file)
	log.Println("I am a log")
}
