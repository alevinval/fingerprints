package kernel

import "github.com/alevinval/fingerprints/internal/matrix"

type multiplication struct {
	Base
	a, b   *matrix.M
	offset int
}

func NewMultiplication(a, b *matrix.M, offset int) *multiplication {
	k := &multiplication{a: a, b: b, offset: offset}
	k.Base = Base{kernel: k}
	return k
}

func (mk *multiplication) Offset() int {
	return mk.offset
}

func (mk *multiplication) Apply(_ *matrix.M, x, y int) float64 {
	sum := 0.0
	for j := -mk.Offset(); j <= mk.Offset(); j++ {
		for i := -mk.Offset(); i <= mk.Offset(); i++ {
			a := mk.a.At(x+i, y+j)
			b := mk.b.At(x+i, y+j)
			sum += a * b
		}
	}
	return sum
}
