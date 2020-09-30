package processing

import (
	"image"
	"math"
	"sync"

	"github.com/alevinval/fingerprints/internal/helpers"
	"github.com/alevinval/fingerprints/internal/matrix"
	"github.com/alevinval/fingerprints/internal/types"
)

func Normalize(in, out *matrix.M, meta types.Metadata) {
	helpers.RunInParallel(in, 0, func(wg *sync.WaitGroup, bounds image.Rectangle) {
		doNormalize(in, out, bounds, meta.MinValue, meta.MaxValue)
		wg.Done()
	})
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
