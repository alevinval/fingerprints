package main

import "math"

type directionalKernel struct {
	gx, gy *Matrix
}

func NewDirectionalKernel(gx, gy *Matrix) *directionalKernel {
	return &directionalKernel{gx: gx, gy: gy}
}

func (k *directionalKernel) Offset() int {
	return 0
}

func (k *directionalKernel) Apply(_ *Matrix, x, y int) float64 {
	dx := k.gx.At(x, y)
	dy := k.gy.At(x, y)
	ang := math.Atan2(dy, dx) + math.Pi/2
	return ang
}
