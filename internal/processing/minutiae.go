package processing

import "github.com/alevinval/fingerprints/internal/matrix"

func ExtractMinutiae(in *matrix.M) {
	bounds := in.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			in.At(x, y)
		}
	}
}
