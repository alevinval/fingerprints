package main

import (
	"math"
)

type fdKernel struct {
	size      int
	gx        *Matrix
	gy        *Matrix
	gxyKernel *mulSumKernel
}

func NewFilteredDirectional(gx, gy *Matrix) *fdKernel {
	return &fdKernel{gx: gx, gy: gy, gxyKernel: NewMulSumKernel(gx, gy)}
}

func (k *fdKernel) Offset() int {
	return Sum8x8.Offset()
}

func (k *fdKernel) Apply(_ *Matrix, x, y int) float64 {
	gxx := Sum8x8.Apply(k.gx, x, y)
	gyy := Sum8x8.Apply(k.gy, x, y)
	gxy := k.gxyKernel.Apply(nil, x, y)
	phy := math.Pi/2 + 0.5*math.Atan2(2*gxy, gxx-gyy)
	return phy
}
