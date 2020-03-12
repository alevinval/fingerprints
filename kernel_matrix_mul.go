package main

type matrixMulKernel struct {
	BaseKernel
	a, b   *Matrix
	offset int
}

func NewKernelMatrixMul(a, b *Matrix, offset int) *matrixMulKernel {
	k := &matrixMulKernel{a: a, b: b, offset: offset}
	k.BaseKernel = BaseKernel{kernel: k}
	return k
}

func (mk *matrixMulKernel) Offset() int {
	return mk.offset
}

func (mk *matrixMulKernel) Apply(_ *Matrix, x, y int) float64 {
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
