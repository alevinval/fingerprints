package kernel

import (
	"math"

	"github.com/alevinval/fingerprints/internal/matrix"
)

type Sqrt struct {
	Base
	a *matrix.M
	b *matrix.M
}

func NewSqrt(a, b *matrix.M) *Sqrt {
	k := &Sqrt{a: a, b: b}
	k.Base = Base{kernel: k}
	return k
}

func (k *Sqrt) Offset() int {
	return 0
}

func (k *Sqrt) Apply(_ *matrix.M, x, y int) float64 {
	aa := k.a.At(x, y)
	bb := k.b.At(x, y)
	return math.Sqrt(aa*aa + bb*bb)
}
