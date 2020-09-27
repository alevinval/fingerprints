package kernel

import (
	"math"

	"github.com/alevinval/fingerprints/internal/matrix"
)

type sqrt struct {
	Base
	a *matrix.M
	b *matrix.M
}

func Sqrt(a, b *matrix.M) *sqrt {
	k := &sqrt{a: a, b: b}
	k.Base = Base{kernel: k}
	return k
}

func (k *sqrt) Offset() int {
	return 0
}

func (k *sqrt) Apply(_ *matrix.M, x, y int) float64 {
	aa := k.a.At(x, y)
	bb := k.b.At(x, y)
	return math.Sqrt(aa*aa + bb*bb)
}
