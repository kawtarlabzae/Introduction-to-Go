package mathUtil

import "math"

func Square(x float64) float64 {
    return x * x
}

func GCD(a, b int) int {
    for b != 0 {
        a, b = b, a%b
    }
    return int(math.Abs(float64(a))) 
}