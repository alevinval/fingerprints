package main

var (
	SobelDx = &sobel{mat: [3][3]float64{{-1, 0, 1}, {-2, 0, 2}, {-1, 0, 1}}}
	SobelDy = &sobel{mat: [3][3]float64{{-1, -2, -1}, {0, 0, 0}, {1, 2, 1}}}
)

func init() {
	SobelDx.BaseKernel = BaseKernel{kernel: SobelDx}
	SobelDy.BaseKernel = BaseKernel{kernel: SobelDy}
}

type sobel struct {
	BaseKernel
	mat [3][3]float64
}

func (k *sobel) Offset() int {
	return 1
}

func (k *sobel) Apply(in *Matrix, x, y int) float64 {
	sum := 0.0
	for j := -k.Offset(); j <= k.Offset(); j++ {
		for i := -k.Offset(); i <= k.Offset(); i++ {
			sum += k.mat[j+1][i+1] * in.At(x+i, y+j)
		}
	}
	return sum
}
