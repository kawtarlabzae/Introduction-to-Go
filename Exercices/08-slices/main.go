package main

import "fmt"

func exampleFunction() {
	// Declare a slice in different ways
	var list1 []int                      // Nil slice (uninitialized)
	list2 := []int{1, 2, 3}              // Literal declaration
	list3 := make([]int, 5)              // Slice with predefined length (defaulted to 0)
	list4 := make([]int, 3, 5)           // Slice with length 3, capacity 5
	list5 := list2[:2]                   // Slice from an existing slice (slicing)
	list6 := append([]int{}, list2...)   // Copy by appending all elements

	// Defer a function to output all lists
	defer func() {
		fmt.Println("\nDeferred Output:")
		fmt.Printf("list1 (nil): %v\n", list1)
		fmt.Printf("list2 (literal): %v\n", list2)
		fmt.Printf("list3 (make with length): %v\n", list3)
		fmt.Printf("list4 (make with length and capacity): %v (capacity: %d)\n", list4, cap(list4))
		fmt.Printf("list5 (slicing): %v\n", list5)
		fmt.Printf("list6 (copy): %v\n", list6)
	}()

	// Modify lists and demonstrate slice methods
	fmt.Println("Function is running...")

	// Append elements
	list1 = append(list1, 10, 20, 30)
	list2 = append(list2, 4)
	
	// Copy elements
	copy(list3, list2)

	// Sub-slicing
	list4 = list2[1:3]

	// Iterating over a slice
	fmt.Println("\nIterating over list2:")
	for i, val := range list2 {
		fmt.Printf("Index %d: %d\n", i, val)
	}

	// Length and capacity
	fmt.Printf("\nLength and capacity of list2: len=%d, cap=%d\n", len(list2), cap(list2))
	fmt.Printf("Length and capacity of list4: len=%d, cap=%d\n", len(list4), cap(list4))

	return
}

func main() {
	exampleFunction()
}
