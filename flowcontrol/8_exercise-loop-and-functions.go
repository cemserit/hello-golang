package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	if x == 0 {
		return 0
	}
	z := 1.0
	for i := 0; i < int(x); i++ {
		z = z - (z*z-x)/(2*z)
	}
	return z

}

func main() {
	for i := 0; i < 100; i++ {
		fmt.Printf("Custom Sqrt(%v)=%g\n", i, Sqrt(float64(i)))
		fmt.Printf("Math Sqrt(%v)=%g\n\n", i, math.Sqrt(float64(i)))
	}
}
