package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	listInt := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, num := range listInt {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			SquareNumber(i)
		}(num)
		
	}
	wg.Wait()
}

func SquareNumber(number int) {
	fmt.Printf("the square number of %d is %d\n", number, number*number)
}
