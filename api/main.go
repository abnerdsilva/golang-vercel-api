package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/api/hello", helloHandler)
	http.ListenAndServe(":8080", nil) // Listen on port 8080
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Set the content type to JSON
	w.Header().Set("Content-Type", "application/json")
	// Respond with a JSON object
	fmt.Fprintf(w, `{"message": "Hello, World from Go on Vercel!"}`)
}
