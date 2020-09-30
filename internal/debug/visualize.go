package debug

import (
	"image"
	"image/color"
	"image/draw"

	"github.com/alevinval/fingerprints/internal/types"
)

var (
	red  = color.RGBA{255, 0, 0, 255}
	blue = color.RGBA{0, 0, 255, 255}
)

func DrawFeatures(original image.Image, result *types.DetectionResult) {
	dst := original.(draw.Image)

	drawFrame(dst, result.Frame.Horizontal, blue)
	drawFrame(dst, result.Frame.Vertical, blue)

	for _, minutiae := range result.Minutia {
		drawFillSquare(dst, minutiae.X, minutiae.Y, red)
	}
}

func drawFrame(dst draw.Image, r image.Rectangle, c color.Color) {
	drawFillSquare(dst, r.Bounds().Min.X, r.Bounds().Min.Y, c)
	drawFillSquare(dst, r.Bounds().Max.X, r.Bounds().Max.Y, c)
}

func drawFillSquare(dst draw.Image, x, y int, c color.Color) {
	dst.Set(x, y, c)
	dst.Set(x+1, y, c)
	dst.Set(x+1, y+1, c)
	dst.Set(x+1, y-1, c)
	dst.Set(x, y-1, c)
	dst.Set(x, y+1, c)
}
