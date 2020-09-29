package processing

import (
	"github.com/alevinval/fingerprints/internal/matrix"
	"github.com/alevinval/fingerprints/internal/types"
)

func ExtractFrame(binarizedSegmented *matrix.M) types.Frame {
	bounds := binarizedSegmented.Bounds()
	xmin, ymin := 9999, 9999
	xmax, ymax := 0, 0
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			v := binarizedSegmented.At(x, y)
			if v > 125 {
				if x < xmin {
					xmin = x
				}
				if y < ymin {
					ymin = y
				}
				if x > xmax {
					xmax = x
				}
				if y > ymax {
					ymin = y
				}

			}
		}
	}

	return types.Frame{xmin, ymin, xmax, ymax}
}
