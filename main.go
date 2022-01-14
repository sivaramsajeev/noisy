package main

import (
	"image/png"
	"math/rand"
	"os"
	"image"
	"image/color"
)

const (
	noOfPixels = 256
)

func main() {

	img := image.NewRGBA(image.Rectangle{
		Min: image.Point{X: 0, Y: 0},
		Max: image.Point{X: noOfPixels, Y: noOfPixels},
	})

	colorate := func() color.RGBA {
		return color.RGBA{R: uint8(rand.Intn(255)), G: uint8(rand.Intn(255)), B: uint8(rand.Intn(255)), A: 0xff}
	}

	for x := 0; x < noOfPixels; x++ {
		for y := 0; y < noOfPixels; y++ {
			img.Set(x, y, colorate())
		}
	}

	f, _ := os.Create("noisy.png")
	_ = png.Encode(f, img)
}
