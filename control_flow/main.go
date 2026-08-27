package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// ============================================================
	// if, else if and else statement
	// ============================================================
	currentTime := time.Now()
	fmt.Println("Current Time:", currentTime)

	if currentHour := time.Now().Hour(); currentHour < 12 {
		fmt.Println("Good morning!")
	} else if currentHour < 17 {
		fmt.Println("Good afternoon!")
	} else {
		fmt.Println("Good evening!")
	}

	// ============================================================
	// switch statement
	// ============================================================
	os := runtime.GOOS
	fmt.Println("Operating System:", os)

	switch my_os := runtime.GOOS; my_os {
	case "windows":
		fmt.Println("Go run -> Windows")
	case "linux":
		fmt.Println("Go run -> Linux")
	case "darwin":
		fmt.Println("Go run -> MacOS")
	default:
		fmt.Println("Go run -> Other OS")
	}

	// ============================================================
	// for loop, break and continue
	// ============================================================
	// counter; condition; update
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		} else if i == 5 {
			break
		}
		fmt.Println(i)
	}

}
