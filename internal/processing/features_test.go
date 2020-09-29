package processing

import (
	"image"
	"testing"

	"github.com/alevinval/fingerprints/internal/matrix"
	"github.com/alevinval/fingerprints/internal/types"
	"github.com/stretchr/testify/assert"
)

func TestExtractFeatures(t *testing.T) {
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

	minutiaes := ExtractFeatures(skeleton, filteredDirectional, segmented)
	assert.NotEmpty(t, minutiaes)
	first := minutiaes[0]
	assert.Equal(t, types.Bifurcation, first.Type)
	assert.Equal(t, 1.0, first.Angle)
	assert.Equal(t, 1, first.X)
	assert.Equal(t, 1, first.Y)
}

func setMatrixTo(m *matrix.M, value float64) {
	bounds := m.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			m.Set(x, y, value)
		}
	}
}
