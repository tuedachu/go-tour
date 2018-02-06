package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println("Iteration ", i, " : z = ", z)
	}
	return z
}

func Sqrt_improved(x float64) float64 {
	z := 1.0
	change := 1.0
	for i := 0; math.Abs(change) > 0.0000001; i++ {
		change = (z*z - x) / (2 * z)
		z -= change
		fmt.Println("Iteration ", i, " : z = ", z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt_improved(2))
}
