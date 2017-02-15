package main

type mulSumKernel struct {
	a, b *Matrix
}

func NewMulSumKernel(a, b *Matrix) *mulSumKernel {
	return &mulSumKernel{a: a, b: b}
}

func (kw *mulSumKernel) Offset() int {
	return 4
}

func (kw *mulSumKernel) Apply(_ *Matrix, x, y int) float64 {
	sum := 0.0
	for i := -kw.Offset(); i <= kw.Offset(); i++ {
		for j := -kw.Offset(); j <= kw.Offset(); j++ {
			a := kw.a.At(x+i, y+j)
			b := kw.b.At(x+i, y+j)
			sum += a * b
		}
	}
	return sum
}
