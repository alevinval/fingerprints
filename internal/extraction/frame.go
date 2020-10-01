package extraction

import (
	"image"
	"log"
	"math"

	"github.com/alevinval/fingerprints/internal/matrix"
	"github.com/alevinval/fingerprints/internal/types"
)

// Frame detects the boundaries of the fingerprint and establishes
// a reference point and the angle of such reference point.
func Frame(binarizedSegmented *matrix.M) types.Frame {
	h := findHorizontalAxis(binarizedSegmented, false)
	v := findVerticalAxis(binarizedSegmented, false)
	d := image.Rect(h.Min.X, v.Min.Y, h.Max.X, v.Max.Y)

	hx, hy := halfPoint(h)
	vx, vy := halfPoint(v)
	angle := math.Sin(float64(hx-vx) / float64(hy-vy))

	log.Printf("frame angle: %f", angle*180/math.Pi)
	return types.Frame{Horizontal: h, Vertical: v, Diagonal: d, Angle: angle}
}

type axis byte

const (
	xAxis axis = iota
	yAxis
)

func findVerticalAxis(binarizedSegmented *matrix.M, isReversed bool) image.Rectangle {
	frame := findAxis(binarizedSegmented, xAxis, yAxis, isReversed)
	log.Printf("vertical frame: %s", frame)
	return frame

}

func findHorizontalAxis(binarizedSegmented *matrix.M, isReversed bool) image.Rectangle {
	frame := findAxis(binarizedSegmented, yAxis, xAxis, false)
	log.Printf("horizontal frame: %s", frame)
	return frame
}

func findAxis(in *matrix.M, firstAxis, secondAxis axis, isReversed bool) image.Rectangle {
	bounds := in.Bounds()
	longestY := 0
	a0, b0, b1 := 0, 0, 0
	providePoints(bounds, firstAxis, isReversed, func(a int) {
		c := 0
		_b0 := 0
		providePoints(bounds, secondAxis, false, func(b int) {
			var v float64
			if firstAxis == xAxis {
				v = in.At(a, b)
			} else {
				v = in.At(b, a)
			}
			if c == 0 && v < 125 {
				// do nothing
			} else if v > 125 && _b0 == 0 {
				_b0 = b
				c++
			} else if v > 125 {
				c++
			} else {
				if c > longestY {
					longestY = c
					a0 = a
					b0 = _b0
					b1 = b
				}
			}
		})
	})

	var frame image.Rectangle
	if firstAxis == xAxis {
		frame = image.Rect(a0, b0, a0, b1)
	} else {
		frame = image.Rect(b0, a0, b1, a0)
	}
	return frame
}

func providePoints(bounds image.Rectangle, ax axis, isReversed bool, f func(n int)) {
	var ini, max int
	if ax == xAxis {
		ini = bounds.Min.X
		max = bounds.Max.X
	} else {
		ini = bounds.Min.Y
		max = bounds.Max.Y
	}

	if isReversed {
		for n := max - 1; n >= ini; n-- {
			f(n)
		}
	} else {
		for n := ini; n < max; n++ {
			f(n)
		}
	}
}

func mergeFrame(a, b types.Frame) types.Frame {
	return types.Frame{
		Horizontal: halfPointAB(a.Horizontal, b.Horizontal),
		Vertical:   halfPointAB(a.Vertical, b.Vertical),
	}
}

func halfPointAB(a, b image.Rectangle) image.Rectangle {
	return image.Rect(
		int((a.Min.X+b.Min.X)/2),
		int((a.Min.Y+b.Min.Y)/2),
		int((a.Max.X+b.Max.X)/2),
		int((a.Max.Y+b.Max.Y)/2),
	)
}

func halfPoint(r image.Rectangle) (int, int) {
	return (r.Max.X + r.Min.X) / 2, (r.Max.Y + r.Min.Y) / 2
}
