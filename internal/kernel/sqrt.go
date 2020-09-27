package kernel

import (
	"math"

	"github.com/alevinval/fingerprints/internal/matrix"
)

type Sqrt struct {
	BaseKernel
	a *matrix.Matrix
	b *matrix.Matrix
}

func NewSqrtKernel(a, b *matrix.Matrix) *Sqrt {
	k := &Sqrt{a: a, b: b}
	k.BaseKernel = BaseKernel{kernel: k}
	return k
}

func (k *Sqrt) Offset() int {
	return 0
}

func (k *Sqrt) Apply(_ *matrix.Matrix, x, y int) float64 {
	aa := k.a.At(x, y)
	bb := k.b.At(x, y)
	return math.Sqrt(aa*aa + bb*bb)
}
