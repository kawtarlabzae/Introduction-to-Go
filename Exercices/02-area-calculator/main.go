package main

import (
	"fmt"

	"introduction-to-go/Exercices/02-area-calculator/calcul"
)

func main() {
	width := 5.0
	length := 4.0
	area := calcul.RectangleArea(length, width)
	fmt.Printf("Area: %.2f square meters\n", area)
	areaConverted := calcul.MetersToFeet(area)
	fmt.Printf("Area in feet: %.2f square feet\n", areaConverted)
}
