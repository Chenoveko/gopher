package main

import (
	"fmt"
	"slices"
)

func main() {
	// ============================================================
	// Arrays
	// ============================================================
	// Fixed-size collection of homogeneous data (data of the same type)
	var numbers = [5]int{1, 2, 3, 4, 5}
	numbers[0] = 0

	var other_numbers = [...]int{45, 23}

	fmt.Println("Number Array: ", numbers)
	fmt.Println("First Element: ", numbers[0])
	fmt.Println("Other Number Array: ", other_numbers)

	for i := 0; i < len(numbers); i++ {
		fmt.Println("Iterate: ", numbers[i])
	}

	for index, value := range numbers {
		fmt.Println("Index: ", index, "| Value: ", value)
	}

	for index, _ := range numbers {
		fmt.Println("Index: ", index)
	}

	var matrix = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println("Matrix 3x3: ", matrix)

	// ============================================================
	// Slice
	// ============================================================
	// Collection of homogeneous data (data of the same type) that can grow or shrink
	daysWeek := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	daysWeekSlice := daysWeek[0:5]
	fmt.Println(daysWeek)
	fmt.Println(daysWeekSlice)

	// Capacity of slice
	fmt.Println("Length of Slice: ", len(daysWeekSlice))
	fmt.Println("Capacity of Slice: ", cap(daysWeekSlice))
	daysWeekSlice = append(daysWeekSlice, "Saturday", "Sunday")

	fmt.Println("After append new elements")
	fmt.Println("Length of Slice: ", len(daysWeekSlice))
	fmt.Println("Capacity of Slice: ", cap(daysWeekSlice))
	daysWeekSlice = append(daysWeekSlice, "Other")

	fmt.Println("After append new element")
	fmt.Println("Length of Slice: ", len(daysWeekSlice))
	fmt.Println("Capacity of Slice: ", cap(daysWeekSlice))

	// Remove elements
	fmt.Println("Before Remove")
	fmt.Println("Slice: ", daysWeekSlice)
	fmt.Println("After Remove")
	daysWeekSlice = slices.Delete(daysWeekSlice, 0, 2)
	fmt.Println("Slice: ", daysWeekSlice)

	// Create empty slice
	empty_slice := make([]string, 5, 10)
	fmt.Println("Empty Slice: ", empty_slice)
	fmt.Println("Length of Slice: ", len(empty_slice))
	fmt.Println("Capacity of Slice: ", cap(empty_slice))

	// Copy slice
	slice_1 := []int{1, 2, 3, 4, 5}
	slice_2 := make([]int, 5)

	number_of_copies := copy(slice_2, slice_1) // Copy values from slice_1 to slice_2
	fmt.Println("Slice 2: ", slice_2)
	fmt.Println("# Copied values: ", number_of_copies)

	// ============================================================
	// Maps
	// ============================================================
	// Key <-> Value
	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#00FF00",
		"blue":  "#0000FF",
	}
	fmt.Println("Colors Map: ", colors)
	fmt.Println("Red Value: ", colors["red"])

	// Add new element
	colors["black"] = "#000000"
	fmt.Println("Colors Map with new element: ", colors)

	// Verification of element
	value, ok := colors["black"]
	fmt.Println("Black Value: ", value, "| Verification: ", ok)

	value, ok = colors["white"]
	fmt.Println("White Value: ", value, "| Verification: ", ok)

	if value, ok := colors["white"]; ok {
		fmt.Println("There is value: ", value)
	} else {
		fmt.Println("There is no value")
	}

	// Delete element
	delete(colors, "blue")
	fmt.Println("Colors Map after delete: ", colors)

	// Iterate map
	for key, value := range colors {
		fmt.Println("Key: ", key, "| Value: ", value)
	}

}
