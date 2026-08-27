package main

import (
	"fmt"
	"math"
	"strconv"
)

// ============================================================
// Constants
// ============================================================
const Pi float32 = 3.14
const (
	X = 100    // Decimal
	Y = 0b1010 // Binary -> 0b
	Z = 0o12   // Octal -> 0o
	W = 0xff   // Hexadecimal -> 0x
)

const (
	Monday = iota + 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
	// ============================================================
	// Variables
	// ============================================================
	// Variable declaration
	var firstName string
	var lastName string
	var age int

	var (
		lunch, dinner    string
		grams_of_protein int
	)

	// Variable Assignment
	firstName = "Kevin"
	lastName = "Thompson"
	age = 83
	lunch = "pizza"
	dinner = "pasta"
	grams_of_protein = 87

	// Variable Assignment and Declaration
	var pokemonName, pokemonType = "Squirtle", "Water"

	// Short Variable Declaration
	apples, oranges := 23, 45

	// Display Variable Values
	fmt.Println("First name:", firstName, "| Last name:", lastName, "| Age:", age)
	fmt.Println("Lunch:", lunch, "| Dinner:", dinner, "| Grams of protein:", grams_of_protein)
	fmt.Println("Pokemon Name:", pokemonName, "| Pokemon Type:", pokemonType)
	fmt.Println("# apples:", apples, "| # oranges:", oranges)
	fmt.Println("Pi:", Pi)
	fmt.Println("X:", X, "| Y:", Y, "| Z:", Z, "| W:", W)
	fmt.Println("Weedk days:", Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)

	// ============================================================
	// Basic Data Types
	// ============================================================
	// Integers (Signed and Unsigned) -> int8, int16, int32, int64, uint8, uint16, uint32, uint64. int and uint depends on system arch
	var my_integer int
	var my_uinteger int
	my_integer = -8
	my_uinteger = 8
	fmt.Println("My signed integer:", my_integer)
	fmt.Println("My unsigned integer:", my_uinteger)
	fmt.Println("Max int64:", math.MaxInt64)
	fmt.Println("Min int64:", math.MinInt64)

	// Floats -> float32 and float64
	var twoRootSquare float32
	twoRootSquare = float32(math.Sqrt(2))
	fmt.Println("Square root of 2:", twoRootSquare)
	fmt.Println("Max float32:", math.MaxFloat32)

	// Booleans
	var is_strong bool
	is_strong = false
	fmt.Println("Is strong:", is_strong)

	// Strings
	fullName := "Robert Griesemer\t\"(alias RG\")\n"
	fmt.Println("Full Name:", fullName)

	// Byte: an alias for uint8
	var a byte = 'a'
	fmt.Println("ASCII code for a:", a)

	// A string is a sequence of bytes
	s := "Hello"
	fmt.Println("ASCII code for H:", s[0])

	// Rune: an alias for int32, used for Unicode characters
	var gopher rune = '🐹'
	fmt.Println("Unicode code point for 🐹:", gopher)

	// ============================================================
	// Default Values
	// ============================================================
	var (
		defaultInt    int
		defaultUint   uint
		defaultFloat  float32
		defaultBool   bool
		defaultString string
	)
	// Display all default values
	fmt.Println("defaultInt:", defaultInt, "defaultUint:", defaultUint, "defaultFloat:", defaultFloat, "defaultBool:", defaultBool, "defaultString:", defaultString)

	// ============================================================
	// Casting
	// ============================================================
	// In Go casting is explicit
	var integer16 int16 = 16
	var integer32 int32 = 50
	fmt.Println(int32(integer16) + integer32)

	number_string := "100"
	number_casted, _ := strconv.Atoi(number_string) // string -> int
	fmt.Println("String -> Integer: ", int32(number_casted+number_casted))

	string_number := 42
	string_number_casted := strconv.Itoa(string_number) // int -> string
	fmt.Println("Integer -> String: ", string_number_casted+string_number_casted)

	// ============================================================
	// fmt Package
	// ============================================================
	fmt.Print("Hello")
	fmt.Print(" darling\n")

	name := "Golang"
	year := 2009

	fmt.Printf("Hello, my name is %s and I was published in %d\n", name, year)
	fmt.Printf("name data type: %T\n", name)
	fmt.Printf("year data type: %T\n", year)

	greeting := fmt.Sprintf("Hello, my name is %s and I was published in %d\n", name, year)
	fmt.Println(greeting)

	var string_input string
	var int_input int
	fmt.Print("String input: ")
	fmt.Scanln(&string_input)
	fmt.Print("Int input: ")
	fmt.Scanln(&int_input)
	fmt.Println("My inputs: ", string_input, int_input)

	// ============================================================
	// Math Operators and the math Package
	// ============================================================
	x := 10.0
	y := 3.0

	// Basic arithmetic operators
	fmt.Println("Addition:", x+y)
	fmt.Println("Subtraction:", x-y)
	fmt.Println("Multiplication:", x*y)
	fmt.Println("Division:", x/y)

	// Functions from the math package
	fmt.Println("Power:", math.Pow(2, 3))
	fmt.Println("Square root:", math.Sqrt(64))
	fmt.Println("Cube root:", math.Cbrt(27))

	// Mathematical constants
	fmt.Println("Pi:", math.Pi)
	fmt.Println("Euler's number:", math.E)

	// Increment and decrement
	x++
	y--
	fmt.Println("x after increment:", x)
	fmt.Println("y after decrement:", y)

	// Assignment operator
	x += 1
	fmt.Println("x after x += 1:", x)
}
