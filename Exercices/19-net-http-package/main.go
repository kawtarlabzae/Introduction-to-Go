package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world")
}
func handler1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello there")
}

func main() {
	http.HandleFunc("/", handler)
	port := ":8080"
	http.HandleFunc("/books", handler1)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
	
}
