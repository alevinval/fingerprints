package processing

import (
	"math"

	"github.com/alevinval/fingerprints/internal/matrix"
	"github.com/alevinval/fingerprints/internal/types"
)

// Metadata about the input image that can be re-used across
// several steps, without having to re-compute it again
func Metadata(in *matrix.M) types.Metadata {
	min, max, mean := findMinMaxMean(in)
	return types.Metadata{
		MinValue:  min,
		MaxValue:  max,
		MeanValue: mean,
	}
}

func findMinMaxMean(in *matrix.M) (float64, float64, float64) {
	var min, max, sum float64
	min = math.MaxFloat64
	bounds := in.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			val := in.At(x, y)
			sum += val
			if val > max {
				max = val
			}
			if val < min {
				min = val
			}
		}
	}
	return min, max, sum / float64(bounds.Dx()*bounds.Dy())
}
