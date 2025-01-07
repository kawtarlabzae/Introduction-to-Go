package main

import "fmt"

func exampleFunction() {
	x := 10
	y := 20

	defer func() {
		fmt.Println("Deferred Output:")
		fmt.Printf("x = %d, y = %d\n", x, y)
	}()

	fmt.Println("Function is running...")
	x += 5
	y += 15
	fmt.Println("Updated values inside the function:", x, y)

	return
}


