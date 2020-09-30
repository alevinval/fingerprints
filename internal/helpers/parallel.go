package helpers

import (
	"image"

	"github.com/alevinval/fingerprints/internal/matrix"
)

// GenerateSubBounds based on an input matrix. Useful to run algorithms
// in parallel, each focusing on a sub-bound.
func GenerateSubBounds(in *matrix.M, offset int) []image.Rectangle {
	subBounds := []image.Rectangle{}
	bounds := in.Bounds()
	blockSize := bounds.Max.X / 2
	for x := bounds.Min.X; x < bounds.Max.X; x += blockSize {
		xi := x - offset
		if xi < bounds.Min.X {
			xi = bounds.Min.X
		}
		xp := x + blockSize + offset
		if xp > bounds.Max.X {
			xp = bounds.Max.X
		}
		for y := bounds.Min.Y; y < bounds.Max.Y; y += blockSize {
			yi := y - offset
			if yi < bounds.Min.Y {
				yi = bounds.Min.Y
			}
			yp := y + blockSize + offset
			if yp > bounds.Max.Y {
				yp = bounds.Max.Y
			}
			subBound := image.Rect(xi, yi, xp, yp)
			subBounds = append(subBounds, subBound)
		}
	}
	return subBounds
}
