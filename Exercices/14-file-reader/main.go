package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name  string
	Age   int
	Email string
}

func main() {
	file, err := os.ReadFile("file-person.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	var people []Person
	err = json.Unmarshal(file, &people)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	for _, person := range people {
		fmt.Println(person)
	}
}
