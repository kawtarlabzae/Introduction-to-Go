// Working with maps in Go
package main

import "fmt"

func main() {
	// 1. Declare and initialize maps
	var map1 map[string]int                   // Nil map, must be initialized before use
	map2 := make(map[string]int)              // Empty map, ready to use
	map3 := map[string]int{"A": 1, "B": 2} // Map with initial values

	// Initialize map1 to make it usable
	map1 = make(map[string]int)

	// 2. Adding elements to a map
	map1["X"] = 100
	map1["Y"] = 200

	map2["Apple"] = 10
	map2["Banana"] = 20

	// 3. Accessing elements in a map
	fmt.Println("Value for key 'X' in map1:", map1["X"])
	fmt.Println("Value for key 'Apple' in map2:", map2["Apple"])

	// 4. Checking if a key exists
	value, exists := map1["Z"]
	if exists {
		fmt.Println("Key 'Z' exists in map1 with value:", value)
	} else {
		fmt.Println("Key 'Z' does not exist in map1")
	}

	// 5. Iterating over a map
	fmt.Println("Iterating over map3:")
	for key, value := range map3 {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}

	// 6. Deleting an element from a map
	delete(map2, "Apple") // Remove the key "Apple"
	fmt.Println("After deleting 'Apple' from map2:", map2)

	// 7. Getting the length of a map
	fmt.Println("Length of map3:", len(map3))

	// 8. Passing a map to a function (maps are reference types)
	updateMap(map3)
	fmt.Println("Map3 after updateMap function:", map3)

	// 9. Nested maps
	nestedMap := make(map[string]map[string]int)
	nestedMap["OuterKey"] = map[string]int{"InnerKey": 42}
	fmt.Println("Nested map value for [OuterKey][InnerKey]:", nestedMap["OuterKey"]["InnerKey"])
}

// Function to update a map
func updateMap(m map[string]int) {
	m["C"] = 3 // Add a new key-value pair
	m["A"] = 10 // Update an existing key
}
