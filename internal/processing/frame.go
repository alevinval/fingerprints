package processing

import (
	"github.com/alevinval/fingerprints/internal/matrix"
)

func ExtractOrigin(binarizedSegmented *matrix.M) (int, int) {
	bounds := binarizedSegmented.Bounds()
	minx, maxy := 9999, 0
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			v := binarizedSegmented.At(x, y)
			if v > 0 {
				if x < minx {
					minx = x
				}
				if y > maxy {
					maxy = y
				}
			}
		}
	}

	return minx, maxy
}
