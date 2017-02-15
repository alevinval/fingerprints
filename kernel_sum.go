package main

var (
	Sum8x8 = &sumKernel{}
)

type sumKernel struct {
	size int
}

func (k *sumKernel) Offset() int {
	return 4
}

func (k *sumKernel) Apply(in *Matrix, x, y int) float64 {
	sum := 0.0
	for i := -k.Offset(); i <= k.Offset(); i++ {
		for j := -k.Offset(); j <= k.Offset(); j++ {
			val := in.At(x+i, y+j)
			sum += val * val
		}
	}
	return sum
}
