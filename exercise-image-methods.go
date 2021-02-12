package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct {
	W, H int
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.W, i.H)
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) At(x, y int) color.Color {
	v := uint8((x + y) / 2)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{100, 100}
	pic.ShowImage(m)
}
