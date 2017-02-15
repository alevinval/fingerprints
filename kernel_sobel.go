package main

var (
	SobelDx = &sobelKernel{size: 3, mat: [3][3]float64{{-1, 0, 1}, {-2, 0, 2}, {-1, 0, 1}}}
	SobelDy = &sobelKernel{size: 3, mat: [3][3]float64{{-1, -2, -1}, {0, 0, 0}, {1, 2, 1}}}
)

type sobelKernel struct {
	size int
	mat  [3][3]float64
}

func (k *sobelKernel) Offset() int {
	return 1
}

func (k *sobelKernel) Apply(in *Matrix, x, y int) float64 {
	sum := 0.0
	for i := -k.Offset(); i <= k.Offset(); i++ {
		for j := -k.Offset(); j <= k.Offset(); j++ {
			sum += k.mat[j+1][i+1] * in.At(x+i, y+j)
		}
	}
	return sum
}