package kernel

import (
	"fmt"
	"image"
	"testing"

	"github.com/alevinval/fingerprints/internal/helpers"
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
	Multiplication(a, b, 1).Convolute(a, out)

	computedBounds := image.Rect(1, 1, 11, 11)
	assertValues(t, out, computedBounds, 18.0)
}

func TestKernelMultiplicationParallelConvolution(t *testing.T) {
	bounds := image.Rect(0, 0, 12, 12)
	a, b := matrix.New(bounds), matrix.New(bounds)
	out := matrix.New(bounds)
	setMatrixTo(a, 1.0)
	setMatrixTo(b, 2.0)
	Multiplication(a, b, 1).ConvoluteParallelized(a, out)

	computedBounds := image.Rect(1, 1, 11, 11)
	assertValues(t, out, computedBounds, 18.0)
}

func TestGenerateSubImagesBounds(t *testing.T) {
	a := newMatrix(0, 0, 12, 8)
	boundsList := helpers.GenerateSubBounds(a, 1)

	expected := []image.Rectangle{
		image.Rect(0, 0, 7, 7),
		image.Rect(0, 5, 7, 8),
		image.Rect(5, 0, 12, 7),
		image.Rect(5, 5, 12, 8),
	}
	assert.Equal(t, expected, boundsList)
}

func newMatrix(a, b, c, d int) *matrix.M {
	return matrix.New(image.Rect(a, b, c, d))
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
