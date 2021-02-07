package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	y := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		x := make([]uint8, dx)
		for l := range x {
			x[l] = uint8((i^l) * 2)
		}
		y[i] = x
	}
	return y
}

func main() {
	pic.Show(Pic)
}
