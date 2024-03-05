package main

import (
	"fmt"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	select {
	case <-time.After(5 * time.Second):
		fmt.Fprintln(w, "Hello, world!")
	case <-ctx.Done():
		fmt.Println("Request cancelled")
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
