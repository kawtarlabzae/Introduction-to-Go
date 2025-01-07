package main

import (
	"errors"
	"fmt"
)

// Define a sentinel error
var ErrDivideByZero = errors.New("cannot divide by zero")

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("divide error: %w", ErrDivideByZero) // Wrap the sentinel error
	}
	return a / b, nil
}

func main() {
	_, err := divide(10, 0)
	if errors.Is(err, ErrDivideByZero) { // Compare with the sentinel error
		fmt.Println("Division failed due to zero denominator",err)
	} else if err != nil { // General error handling
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Division succeeded")
	}
}
