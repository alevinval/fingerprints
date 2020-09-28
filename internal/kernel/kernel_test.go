package kernel

import (
	"fmt"
	"image"
	"testing"

	"github.com/alevinval/fingerprints/internal/matrix"
	"github.com/stretchr/testify/assert"
)

/**
Multiplication kernel does (x, y) = SUM((ax, ay) * (bx, by))
The following operation with offset 1:

	1 1 1   2 2 2
	1 1 1 x 2 2 2
	1 1 1   2 2 2

Results in

	0  0 0
	0 18 0
	0  0 0
*/
func TestKernelMultiplicationConvolution(t *testing.T) {
	bounds := image.Rect(0, 0, 12, 12)
	a, b := matrix.New(bounds), matrix.New(bounds)
	out := matrix.New(bounds)
	setMatrixTo(a, 1.0)
	setMatrixTo(b, 2.0)
	Multiplication(a, b, 1).Convolution(a, out)

	computedBounds := image.Rect(1, 1, 11, 11)
	assertValues(t, out, computedBounds, 18.0)
}

func TestKernelMultiplicationParallelConvolution(t *testing.T) {
	bounds := image.Rect(0, 0, 12, 12)
	a, b := matrix.New(bounds), matrix.New(bounds)
	out := matrix.New(bounds)
	setMatrixTo(a, 1.0)
	setMatrixTo(b, 2.0)
	Multiplication(a, b, 1).ParallelConvolution(a, out)

	computedBounds := image.Rect(1, 1, 11, 11)
	assertValues(t, out, computedBounds, 18.0)
}

func setMatrixTo(m *matrix.M, value float64) {
	bounds := m.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			m.Set(x, y, value)
		}
	}
}

func assertValues(t *testing.T, m *matrix.M, bounds image.Rectangle, expected float64) {
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			assert.Equal(t, expected, m.At(x, y), fmt.Sprintf("value different than expected at (%d,%d)", x, y))
		}
	}
}
