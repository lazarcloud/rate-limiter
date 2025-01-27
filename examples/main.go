package main

import (
	"fmt"
	"net/http"
	"time"

	limiter "github.com/lazarcloud/rate-limiter"
)

func main() {
	// Create a new rate limiter (e.g., max 10 requests per minute per client).
	rl := limiter.New(10, time.Minute)

	// Sample HTTP handler.
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, world!")
	})

	// Wrap the handler with the rate limiter middleware.
	http.Handle("/", rl.Middleware(handler))

	// Start the server.
	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", nil)
}
