package processing

import (
	"image"
	"log"

	"github.com/alevinval/fingerprints/internal/matrix"
	"github.com/alevinval/fingerprints/internal/types"
)

func ExtractFrame(binarizedSegmented *matrix.M) types.Frame {
	v := findVerticalAxis(binarizedSegmented)
	h := findHorizontalAxis(binarizedSegmented)
	return types.Frame{Horizontal: h, Vertical: v}
}

func findVerticalAxis(binarizedSegmented *matrix.M) image.Rectangle {
	bounds := binarizedSegmented.Bounds()
	longestX := 0
	x1, y0, y1 := 0, 0, 0
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		c := 0
		_y0 := 0
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			v := binarizedSegmented.At(x, y)
			if c == 0 && v < 125 {
				continue
			} else if v > 125 && _y0 == 0 {
				_y0 = y
				c++
			} else if v > 125 {
				c++
			} else {
				if c > longestX {
					longestX = c
					y0 = _y0
					x1 = x
					y1 = y
				}
			}
		}
	}

	frame := image.Rect(x1, y0, x1, y1)
	log.Printf("vertical frame: %s", frame)
	return frame
}

func findHorizontalAxis(binarizedSegmented *matrix.M) image.Rectangle {
	bounds := binarizedSegmented.Bounds()
	longestY := 0
	x0, x1, y1 := 0, 0, 0
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		c := 0
		_x0 := 0
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			v := binarizedSegmented.At(x, y)
			if c == 0 && v < 125 {
				continue
			} else if v > 125 && _x0 == 0 {
				_x0 = x
				c++
			} else if v > 125 {
				c++
			} else {
				if c > longestY {
					longestY = c
					x0 = _x0
					x1 = x
					y1 = y
				}
			}
		}
	}

	frame := image.Rect(x0, y1, x1, y1)
	log.Printf("horizontal frame: %s", frame)
	return frame
}
