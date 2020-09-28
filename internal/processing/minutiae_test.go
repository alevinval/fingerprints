package processing

import (
	"image"
	"testing"

	"github.com/alevinval/fingerprints/internal/matrix"
	"github.com/stretchr/testify/assert"
)

func TestExtractMinutiae(t *testing.T) {
	filteredDirectional := matrix.New(image.Rect(0, 0, 3, 3))
	setMatrixTo(filteredDirectional, 1.0)
	segmented := matrix.New(image.Rect(0, 0, 3, 3))
	setMatrixTo(segmented, 1.0)

	// Bifurcation
	skeleton := matrix.New(image.Rect(0, 0, 3, 3))
	skeleton.Set(0, 0, 1.0)
	skeleton.Set(1, 1, 1.0)
	skeleton.Set(1, 2, 1.0)
	skeleton.Set(2, 1, 1.0)

	minutiaes := ExtractMinutiae(skeleton, filteredDirectional, segmented)
	assert.NotEmpty(t, minutiaes)
	first := minutiaes[0]
	assert.Equal(t, Bifurcation, first.Type)
}

func setMatrixTo(m *matrix.M, value float64) {
	bounds := m.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			m.Set(x, y, value)
		}
	}
}
