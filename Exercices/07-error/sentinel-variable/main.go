package main

import (
	"errors"
	"fmt"
)

var ErrDivideByZero = errors.New("cannot divide by zero") // Sentinel error

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}
	return a / b, nil
}

func main() {
	_, err := divide(10, 0)
	if err == ErrDivideByZero { // Compare sentinel error directly
		fmt.Println("Cannot divide by zero!")
	} else if err != nil {
		fmt.Println("Unexpected error:", err)
	} else {
		fmt.Println("Division successful")
	}
}
