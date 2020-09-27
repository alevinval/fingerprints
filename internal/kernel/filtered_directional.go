package kernel

import (
	"math"

	"github.com/alevinval/fingerprints/internal/matrix"
)

type filteredDirectional struct {
	Base
	mulGx, mulGy, mulGxy *multiplication
	offset               int
}

func FilteredDirectional(gx, gy *matrix.M, offset int) *filteredDirectional {
	k := &filteredDirectional{
		mulGx:  Multiplication(gx, gx, offset),
		mulGy:  Multiplication(gy, gy, offset),
		mulGxy: Multiplication(gx, gy, offset),
		offset: offset}
	k.Base = Base{kernel: k}
	return k
}

func (k *filteredDirectional) Offset() int {
	return k.offset
}

func (k *filteredDirectional) Apply(_ *matrix.M, x, y int) float64 {
	gxx := k.mulGx.Apply(nil, x, y)
	gyy := k.mulGy.Apply(nil, x, y)
	gxy := k.mulGxy.Apply(nil, x, y)
	phy := math.Pi/2 + 0.5*math.Atan2(2*gxy, gxx-gyy)
	return phy
}
