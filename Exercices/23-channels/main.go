package main

import (
	"fmt"
	"time"
)

func main() {
	v := 1
	ch := make(chan int)
	go func() {
		s := SquareNumber(v)
		ch <- s
		fmt.Println("hello")
	}()
	vSquare := <- ch
	time.Sleep(2*time.Second)
	fmt.Println(vSquare)
}
func SquareNumber(number int) int{
	return number*number
}