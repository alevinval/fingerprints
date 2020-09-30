package processing

import (
	"image"
	"math"
	"sync"

	"github.com/alevinval/fingerprints/internal/helpers"
	"github.com/alevinval/fingerprints/internal/matrix"
)

func Normalize(in, out *matrix.M) {
	min, max := findMinMax(in)
	helpers.RunInParallel(in, 0, func(wg *sync.WaitGroup, bounds image.Rectangle) {
		doNormalize(in, out, bounds, min, max)
		wg.Done()
	})
}

func findMinMax(in *matrix.M) (float64, float64) {
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
	return min, max
}

func doNormalize(in, out *matrix.M, bounds image.Rectangle, min, max float64) {
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			pixel := in.At(x, y)
			normalizedPixel := math.MaxUint8 * (pixel - min) / (max - min)
			out.Set(x, y, normalizedPixel)
		}
	}
}
