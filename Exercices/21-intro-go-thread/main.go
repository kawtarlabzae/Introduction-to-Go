package main

import (
	"fmt"
	"time"
)

var counter int // Shared variable

func increment() {
	for i := 0; i < 1000; i++ {
		counter++ // Incrementing shared variable
	}
}

func main() {
	go increment() 
	go increment() 


	time.Sleep( time.Millisecond)

	fmt.Println("Final counter value:", counter) 
}
