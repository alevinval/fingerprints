package processing

import (
	"math"

	"github.com/alevinval/fingerprints/internal/matrix"
)

func Normalize(in, out *matrix.M) {
	var min, max float64
	min = math.MaxFloat64

	bounds := in.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			val := in.At(x, y)
			if val > max {
				max = val
			}
			if val < min {
				min = val
			}
		}
	}
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			pixel := in.At(x, y)
			normalizedPixel := math.MaxUint8 * (pixel - min) / (max - min)
			out.Set(x, y, normalizedPixel)
		}
	}
}
