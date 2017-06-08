package example

import (
	"image"
	"image/color"
)

// Gradient returns an image of the given size that contains a white to black
// gradient from top to bottom.
func Gradient(width, height int) image.Image {
	img := image.NewGray(image.Rectangle{Max: image.Point{X: width, Y: height}})
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.SetGray(x, y, color.Gray{Y: 255 - uint8((float64(y)/float64(height))*255)})
		}
	}
	return img
}

// RedOnly returns a version of in with only the color red remaining.
func RedOnly(in image.Image) image.Image {
	b := in.Bounds()
	out := image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
	for x := 0; x < b.Dx(); x++ {
		for y := 0; y < b.Dy(); y++ {
			r, _, _, _ := in.At(x, y).RGBA()
			out.Set(x, y, color.RGBA{R: uint8(r >> 8), A: 255})
		}
	}
	return out
}
