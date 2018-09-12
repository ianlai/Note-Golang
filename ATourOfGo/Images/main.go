package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	rect image.Rectangle
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
	return img.rect
}

func (img Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x), uint8(x + y), uint8(y), 255}
}

func main() {
	m := Image{image.Rect(0, 0, 500, 300)}
	pic.ShowImage(m)
}
