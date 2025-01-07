package main

import "fmt"

func makeMultiplier(multiplier int) func(int) (int, int) {
	i:=0
    return func(value int) (int, int) {
		i++
		multiplier++
        return value * multiplier,i
    }
}

func main() {
    double := makeMultiplier(2) 
    triple := makeMultiplier(3) 

    fmt.Println(double(5)) 
	fmt.Println(double(5))
    fmt.Println(triple(5)) 
}
