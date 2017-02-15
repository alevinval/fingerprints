package main

import "math"

type Sqrt struct {
	a *Matrix
	b *Matrix
}

func NewSqrtKernel(a, b *Matrix) *Sqrt {
	return &Sqrt{a: a, b: b}
}

func (k *Sqrt) Offset() int {
	return 0
}

func (k *Sqrt) Apply(_ *Matrix, x, y int) float64 {
	aa := k.a.At(x, y)
	bb := k.b.At(x, y)
	return math.Sqrt(aa*aa + bb*bb)
}
