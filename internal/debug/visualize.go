package debug

import (
	"image"
	"image/color"
	"image/draw"

	"github.com/alevinval/fingerprints/internal/types"
)

var (
	red = color.RGBA{255, 0, 0, 255}
)

func DrawFeatures(original image.Image, minutia types.MinutiaeList) {
	dst := original.(draw.Image)
	for _, minutiae := range minutia {
		dst.Set(minutiae.X, minutiae.Y, red)
		dst.Set(minutiae.X+1, minutiae.Y, red)
		dst.Set(minutiae.X+1, minutiae.Y+1, red)
		dst.Set(minutiae.X+1, minutiae.Y-1, red)
		dst.Set(minutiae.X, minutiae.Y-1, red)
		dst.Set(minutiae.X, minutiae.Y+1, red)
	}
}
