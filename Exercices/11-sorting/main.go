package main

import "fmt"

func sortAscending(slice []int) {
	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-i-1; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
}

func sortDescending(slice []int) {
	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-i-1; j++ {
			if slice[j] < slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
}

func main() {
	slice := []int{5, 2, 9, 1, 5, 6}
	fmt.Println("Original slice:", slice)

	sortAscending(slice)
	fmt.Println("Sorted in ascending order:", slice)

	sortDescending(slice)
	fmt.Println("Sorted in descending order:", slice)
}
