package main

import "fmt"

func main() {
	// 1. Declare and initialize fixed-size arrays
	var arr1 [5]int                        // Uninitialized array of size 5 (default values are 0)
	arr2 := [5]int{1, 2, 3, 4, 5}          // Fully initialized array
	//arr3 := [5]int{1, 2}                   // Partially initialized array (remaining values are 0)
	arr4 := [...]int{10, 20, 30, 40, 50}   // Compiler infers the size (length = 5)

	// 2. Access and modify elements
	arr1[0] = 100
	arr1[4] = 500
	fmt.Println("Modified arr1:", arr1)

	// 3. Iterate over an array
	fmt.Println("Elements in arr2:")
	for i, v := range arr2 {
		fmt.Printf("Index %d: %d\n", i, v)
	}

	// 4. Copy an array
	arrCopy := arr2
	arrCopy[0] = 99 // Changes only the copy, not the original
	fmt.Println("Original arr2:", arr2)
	fmt.Println("Copied and modified arrCopy:", arrCopy)

	// 5. Compare arrays (only works if arrays are of the same type and size)
	areEqual := arr2 == [5]int{1, 2, 3, 4, 5}
	fmt.Println("Are arr2 and [1, 2, 3, 4, 5] equal?:", areEqual)

	// 6. Multi-dimensional arrays
	var matrix [2][3]int
	matrix[0][0] = 1
	matrix[1][2] = 5
	fmt.Println("Matrix:", matrix)

	// Iterate over a multi-dimensional array
	fmt.Println("Matrix elements:")
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("matrix[%d][%d] = %d\n", i, j, matrix[i][j])
		}
	}

	// 7. Find the length of an array
	fmt.Println("Length of arr4:", len(arr4))

	// 8. Passing arrays to functions
	fmt.Println("Sum of elements in arr2:", sumArray(arr2))
}

// Function to calculate the sum of array elements
func sumArray(arr [5]int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}
