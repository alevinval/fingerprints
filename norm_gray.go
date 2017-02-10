package main

import (
	"image"
	"math"
)

func normalizeGray(in *image.Gray) {
	var min, max uint8
	min = math.MaxUint8
	for x := 0; x < in.Rect.Dx(); x++ {
		for y := 0; y < in.Rect.Dy(); y++ {
			val := in.GrayAt(x, y).Y
			if val < min {
				min = val
			}
			if val > max {
				max = val
			}
		}
	}
	for x := 0; x < in.Rect.Dx(); x++ {
		for y := 0; y < in.Rect.Dy(); y++ {
			c := in.GrayAt(x, y)
			c.Y = uint8(255 * float64(c.Y-min) / float64(max-min))
			in.SetGray(x, y, c)
		}
	}
}
