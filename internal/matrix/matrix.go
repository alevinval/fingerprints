package matrix

import (
	"image"
	"image/color"
)

type M struct {
	pixels [][]float64
	bounds image.Rectangle
}

func New(bounds image.Rectangle) *M {
	dx, dy := bounds.Dx(), bounds.Dy()
	if dx != dy {
		panic("only squared images are supported for the moment")
	}
	picture := make([][]float64, dx)
	pixels := make([]float64, dx*dx)
	for i := range picture {
		picture[i], pixels = pixels[:dx], pixels[dx:]
	}

	m := new(M)
	m.bounds = bounds
	m.pixels = picture
	return m
}

func NewFromGray(in *image.Gray) *M {
	bounds := in.Bounds()
	m := New(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			m.Set(x, y, float64(in.GrayAt(x, y).Y))
		}
	}
	return m
}

func (m *M) At(x, y int) float64 {
	return m.pixels[y][x]
}

func (m *M) Set(x, y int, value float64) {
	m.pixels[y][x] = value
}

func (m *M) Bounds() image.Rectangle {
	return m.bounds
}

func (m *M) SubImage(r image.Rectangle) *M {
	r = r.Intersect(m.bounds)

	// If r1 and r2 are Rectangles, r1.Intersect(r2) is not guaranteed to be inside
	// either r1 or r2 if the intersection is empty. Without explicitly checking for
	// this, the Pix[i:] expression below can panic.
	if r.Empty() {
		panic("wtf")
	}
	return &M{
		pixels: m.pixels,
		bounds: r,
	}
}

func (m *M) ToGray() *image.Gray {
	bounds := m.Bounds()
	gray := image.NewGray(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			gray.SetGray(x, y, color.Gray{Y: uint8(m.At(x, y))})
		}
	}
	return gray
}
