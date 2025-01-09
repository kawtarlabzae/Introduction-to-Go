package main

import "fmt"

func sender(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i
		fmt.Print(i)
	}
	close(ch)
}

func main() {
	ch := make(chan int)

	go sender(ch)

	for {
		val, ok := <-ch
		if !ok {
			fmt.Println("Channel closed!")
			break
		}
		fmt.Println("Received:", val)
	}

	fmt.Println("Done!")
}
