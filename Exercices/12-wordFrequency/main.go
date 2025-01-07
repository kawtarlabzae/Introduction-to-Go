package main

import (
	"fmt"
)

func countOccurrences(names []string) map[string]int {
	map1 := make(map[string]int)

	for _, val := range names {
		_, exists := map1[val]
		if exists {
			map1[val]++
		} else {
			map1[val] = 1
		}
	}

	return map1
}

func main() {
	var n int
	fmt.Println("Enter the number of names:")
	fmt.Scan(&n) 

	names := make([]string, n)
	fmt.Printf("Enter %d names separated by space:\n", n)
	for i := 0; i < n; i++ {
		fmt.Scan(&names[i])
	}


	result := countOccurrences(names)

	for key, value := range result {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}
}
