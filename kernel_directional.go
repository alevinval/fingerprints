package main

import (
	"image"
	"math"
)

type directionalKernel struct {
	gx *image.Gray
	gy *image.Gray
}

func NewDirectionalKernel(gx, gy *image.Gray) *directionalKernel {
	return &directionalKernel{gx: gx, gy: gy}
}

func (k *directionalKernel) Offset() int {
	return 4
}

func (k *directionalKernel) Apply(in *image.Gray, x, y int) float64 {
	dx := k.gx.GrayAt(x, y).Y
	dy := k.gy.GrayAt(x, y).Y
	ang := math.Atan2(float64(dy), float64(dx))
	val := (ang + math.Pi/2) / math.Pi
	return val
}
