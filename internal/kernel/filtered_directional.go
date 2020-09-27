package kernel

import (
	"math"

	"github.com/alevinval/fingerprints/internal/matrix"
)

type filteredDirectional struct {
	BaseKernel
	mulGx, mulGy, mulGxy *matrixMulKernel
	offset               int
}

func NewFilteredDirectional(gx, gy *matrix.Matrix, offset int) *filteredDirectional {
	k := &filteredDirectional{
		mulGx:  NewKernelMatrixMul(gx, gx, offset),
		mulGy:  NewKernelMatrixMul(gy, gy, offset),
		mulGxy: NewKernelMatrixMul(gx, gy, offset),
		offset: offset}
	k.BaseKernel = BaseKernel{kernel: k}
	return k
}

func (k *filteredDirectional) Offset() int {
	return k.offset
}

func (k *filteredDirectional) Apply(_ *matrix.Matrix, x, y int) float64 {
	gxx := k.mulGx.Apply(nil, x, y)
	gyy := k.mulGy.Apply(nil, x, y)
	gxy := k.mulGxy.Apply(nil, x, y)
	phy := math.Pi/2 + 0.5*math.Atan2(2*gxy, gxx-gyy)
	return phy
}
