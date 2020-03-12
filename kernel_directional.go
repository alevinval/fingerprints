package main

import "math"

type directionalKernel struct {
	BaseKernel
	gx, gy *Matrix
}

func NewDirectionalKernel(gx, gy *Matrix) *directionalKernel {
	k := &directionalKernel{gx: gx, gy: gy}
	k.BaseKernel = BaseKernel{kernel: k}
	return k
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
