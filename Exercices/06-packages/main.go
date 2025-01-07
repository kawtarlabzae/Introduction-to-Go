package main

import (
	"introduction-to-go/Exercices/02-area-calculator/calcul"
	"log"
	"um6p/hello/mathUtil"
)

func main() {
	log.Println("hello")
	a, b := 5, 2
	log.Println(calcul.RectangleArea(float64(a), float64(b)))
	log.Println(mathUtil.GCD(a, b))
}
