package main

import "math"

func Normalize(in, out *Matrix) {
	var min, max float64
	min = math.MaxFloat64

	bounds := in.Bounds()
	dx, dy := bounds.Dx(), bounds.Dy()
	for x := 0; x < dx; x++ {
		for y := 0; y < dy; y++ {
			val := in.At(x, y)
			if val > max {
				max = val
			}
			if val < min {
				min = val
			}
		}
	}
	for x := 0; x < dx; x++ {
		for y := 0; y < dy; y++ {
			pixel := in.At(x, y)
			normalizedPixel := math.MaxUint8 * (pixel - min) / (max - min)
			out.Set(x, y, normalizedPixel)
		}
	}
}
