package image

import (
	"image"
)

func ToRGBA(img image.Image) *image.RGBA {
	rgba, ok := img.(*image.RGBA)
	if ok {
		return rgba
	}

	rgba = image.NewRGBA(img.Bounds())
	for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
		for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
			rgba.Set(x, y, img.At(x, y))
		}
	}

	return rgba
}
