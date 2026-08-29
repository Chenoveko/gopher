package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!\n")
	fmt.Fprintf(w, "HTTP Method: %s\n", r.Method)
	fmt.Fprintf(w, "URL Path: %s\n", r.URL.Path)
	fmt.Fprintf(w, "Host: %s\n", r.Host)
	fmt.Fprintf(w, "User Agent: %s\n", r.UserAgent())
}

func api(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"message":"Hello",
		"status":"200",
	}

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`
		<!DOCTYPE html>
		<html>
		<head>
			<title>Health Check</title>
		</head>
		<body>
			<h1>API is healthy!</h1>
			<p>Status: OK</p>
		</body>
		</html>
	`))
}

func greetings(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	w.Write(fmt.Appendf(nil, "Hello, %s!", name))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", helloWorld)
	mux.HandleFunc("GET /api", api)
	mux.HandleFunc("GET /health", healthCheck)
	mux.HandleFunc("GET /greetings/{name}", greetings)

	fmt.Println("Server Running in: http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}
