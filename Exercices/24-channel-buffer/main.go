package main

import "fmt"

func main() {

	ch := make(chan int, 3)
	ch <- 1
	fmt.Print(1)
	ch <- 2
	fmt.Print("he")
	ch <- 3
	fmt.Print("he")
	fmt.Println(<-ch)
	ch <- 4
	fmt.Print("he")

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)

}