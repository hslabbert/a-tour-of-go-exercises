package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		err := ErrNegativeSqrt(x)
		var f float64
		return f, err
	}
	var err error
	z := float64(1)
	var zc []float64
	for i := 0; i < 10; i++ {
		fmt.Println(z)
		z -= (z*z - x) / (2 * z)
		if contains(zc, z) {
			return z, err
		}
		zc = append(zc, z)
	}
	return z, err
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
	var nums = []float64{2, -2}
	for _, v := range nums {
		f, err := Sqrt(v)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(f)
	}
}
