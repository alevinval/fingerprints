package main

import (
	"image"
	"image/color"
)

type Matrix struct {
	pixels [][]float64
	bounds image.Rectangle
}

func NewMatrix(bounds image.Rectangle) *Matrix {
	dx, dy := bounds.Dx(), bounds.Dy()
	picture := make([][]float64, dy)
	pixels := make([]float64, dx*dy)
	for i := range picture {
		picture[i], pixels = pixels[:dx], pixels[dx:]
	}

	m := new(Matrix)
	m.bounds = bounds
	m.pixels = picture
	return m
}

func NewMatrixFromGray(in *image.Gray) *Matrix {
	bounds := in.Bounds()
	m := NewMatrix(bounds)
	for x := 0; x < bounds.Dx(); x++ {
		for y := 0; y < bounds.Dy(); y++ {
			m.Set(x, y, float64(in.GrayAt(x, y).Y))
		}
	}
	return m
}

func (m *Matrix) At(x, y int) float64 {
	return m.pixels[x][y]
}

func (m *Matrix) Set(x, y int, value float64) {
	m.pixels[x][y] = value
}

func (m *Matrix) Bounds() image.Rectangle {
	return m.bounds
}

func (m *Matrix) ToGray() *image.Gray {
	bounds := m.Bounds()
	gray := image.NewGray(bounds)
	for x := 0; x < bounds.Dx(); x++ {
		for y := 0; y < bounds.Dy(); y++ {
			gray.SetGray(x, y, color.Gray{Y: uint8(m.At(x, y))})
		}
	}
	return gray
}
