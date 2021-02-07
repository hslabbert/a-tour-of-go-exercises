package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	var zc []float64
	for i := 0; i < 10; i++ {
		fmt.Println(z)
		z -= (z*z - x) / (2 * z)
		if contains(zc, z) {
			return z
		}
		zc = append(zc, z)
	}
	return z
}

func contains(s []float64, f float64) bool {
	for _, e := range s {
		if e == f {
			return true
		}
	}
	return false
}

func main() {
	num := float64(2)
	fmt.Println(Sqrt(num))
	fmt.Println(math.Sqrt(num))
}
