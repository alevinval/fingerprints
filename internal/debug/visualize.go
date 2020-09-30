package debug

import (
	"image"
	"image/color"
	"image/draw"

	"github.com/alevinval/fingerprints/internal/types"
)

var (
	red   = color.RGBA{255, 0, 0, 255}
	green = color.RGBA{0, 255, 0, 255}
	cyan  = color.RGBA{0, 255, 255, 255}
	blue  = color.RGBA{0, 0, 255, 255}
)

// DrawFeatures draws the original image with all the features that are
// detected drawn on top of it. Useful for understanding what data
// we are gathering, and visualise it. Helpful for detecting issues with
// the algorithms or potential next steps.
func DrawFeatures(original image.Image, result *types.DetectionResult) {
	dst := original.(draw.Image)

	for _, minutiae := range result.Minutia {
		drawSquare(dst, minutiae.X, minutiae.Y, red)
	}

	drawFrame(dst, result.Frame.Horizontal, blue)
	drawFrame(dst, result.Frame.Vertical, blue)
	drawDiagonalFrame(dst, result.Frame.Diagonal, blue)

	drawHalfPoint(dst, result.Frame.Diagonal, cyan)
	drawHalfPoint(dst, result.Frame.Horizontal, red)
	drawHalfPoint(dst, result.Frame.Vertical, green)

}

func drawFrame(dst draw.Image, r image.Rectangle, c color.Color) {
	drawCross(dst, r.Bounds().Min.X, r.Bounds().Min.Y, c)
	drawCross(dst, r.Bounds().Max.X, r.Bounds().Max.Y, c)
}

func drawDiagonalFrame(dst draw.Image, r image.Rectangle, c color.Color) {
	drawEdgeTopLeft(dst, r.Bounds().Min.X, r.Bounds().Min.Y, c)
	drawEdgeBottomRight(dst, r.Bounds().Max.X, r.Bounds().Max.Y, c)
}

func drawHalfPoint(dst draw.Image, r image.Rectangle, c color.Color) {
	halfX, halfY := halfPoint(r)
	drawX(dst, halfX, halfY, c)
}

func drawSquare(dst draw.Image, x, y int, c color.Color) {
	dst.Set(x, y-1, c)
	dst.Set(x, y+1, c)
	dst.Set(x+1, y, c)
	dst.Set(x+1, y-1, c)
	dst.Set(x+1, y+1, c)
	dst.Set(x-1, y, c)
	dst.Set(x-1, y-1, c)
	dst.Set(x-1, y+1, c)
}

func drawCross(dst draw.Image, x, y int, c color.Color) {
	dst.Set(x, y, c)
	dst.Set(x, y-1, c)
	dst.Set(x, y+1, c)
	dst.Set(x+1, y, c)
	dst.Set(x-1, y, c)
}

func drawX(dst draw.Image, x, y int, c color.Color) {
	dst.Set(x, y, c)
	dst.Set(x-1, y-1, c)
	dst.Set(x+1, y+1, c)
	dst.Set(x-1, y+1, c)
	dst.Set(x+1, y-1, c)
}

func drawEdgeTopLeft(dst draw.Image, x, y int, c color.Color) {
	dst.Set(x-1, y+1, c)
	dst.Set(x-1, y, c)
	dst.Set(x-1, y-1, c)
	dst.Set(x, y-1, c)
	dst.Set(x+1, y-1, c)
}

func drawEdgeBottomRight(dst draw.Image, x, y int, c color.Color) {
	dst.Set(x-1, y+1, c)
	dst.Set(x, y+1, c)
	dst.Set(x+1, y+1, c)
	dst.Set(x+1, y, c)
	dst.Set(x+1, y-1, c)
}

func halfPoint(r image.Rectangle) (int, int) {
	return (r.Max.X + r.Min.X) / 2, (r.Max.Y + r.Min.Y) / 2
}
