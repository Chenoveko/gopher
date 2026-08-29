package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// ============================================================
	// Concurrency with API's
	// ============================================================
	// Test API's without concurrency
	start := time.Now()
	apis := [...]string{
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://outlook.office.com",
		"https://api.somewhreintheinternet.com",
		"https://graph.microsoft.com",
	}
	for _, api := range apis {
		checkApi(api)
	}
	elapsed := time.Since(start)
	fmt.Printf("Elapsed time without concurrency -> %v seconds\n", elapsed.Seconds())
	// Test API with go routines and channels
	start = time.Now()
	channel := make(chan string) // We use channels to communicate with go routines
	for _, api := range apis {
		go checkApiConcurrency(api, channel)
	}

	for i := 0; i < len(apis); i++ {
		fmt.Println("Channel: ", <-channel)
	}
	elapsed = time.Since(start)
	fmt.Printf("Elapsed time with go routines 🐹 -> %v seconds\n", elapsed.Seconds())

}

func checkApi(api string) {
	if _, err := http.Get(api); err != nil {
		fmt.Printf("API ERROR: %s is down \n", api)
		return
	}
	fmt.Printf("API %s is running!\n", api)
}

func checkApiConcurrency(api string, channel chan string) {
	if _, err := http.Get(api); err != nil {
		channel <- fmt.Sprintf("API ERROR: %s is down \n", api)
		return
	}
	channel <- fmt.Sprintf("API %s is running!\n", api)
}
