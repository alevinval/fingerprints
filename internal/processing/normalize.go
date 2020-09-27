package processing

import (
	"math"

	"github.com/alevinval/fingerprints/internal/matrix"
)

func Normalize(in, out *matrix.M) {
	var min, max float64
	min = math.MaxFloat64

	bounds := in.Bounds()
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			val := in.At(x, y)
			if val > max {
				max = val
			}
			if val < min {
				min = val
			}
		}
	}
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			pixel := in.At(x, y)
			normalizedPixel := math.MaxUint8 * (pixel - min) / (max - min)
			out.Set(x, y, normalizedPixel)
		}
	}
}
