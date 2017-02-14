package main

import (
	"image"
	"math"
)

type fdKernel struct {
	size int
	gx   *image.Gray
	gy   *image.Gray
}

func NewFilteredDirectional(gx, gy *image.Gray) *fdKernel {
	return &fdKernel{size: 9, gx: gx, gy: gy}
}

func (k *fdKernel) Offset() int {
	return 4
}

func (k *fdKernel) Apply(in *image.Gray, x, y int) float64 {
	dx := Sum9x9.Apply(k.gx, x, y)
	dy := Sum9x9.Apply(k.gy, x, y)
	val := math.Pi/2 + 0.5*math.Atan2(float64(2*dx*dx), float64(dx*dx-dy*dy))
	return val
}
