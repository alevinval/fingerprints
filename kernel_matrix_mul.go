package main

type matrixMulKernel struct {
	a, b   *Matrix
	offset int
}

func NewKernelMatrixMul(a, b *Matrix, offset int) *matrixMulKernel {
	return &matrixMulKernel{a: a, b: b, offset: offset}
}

func (mk *matrixMulKernel) Offset() int {
	return mk.offset
}

func (mk *matrixMulKernel) Apply(_ *Matrix, x, y int) float64 {
	sum := 0.0
	for i := -mk.Offset(); i <= mk.Offset(); i++ {
		for j := -mk.Offset(); j <= mk.Offset(); j++ {
			a := mk.a.At(x+i, y+j)
			b := mk.b.At(x+i, y+j)
			sum += a * b
		}
	}
	return sum
}
