package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("divide error: %v", errors.New("cannot divide by zero"))
	}
	return a / b, nil
}

func main() {
	_, err := divide(10, 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("hi")
}
