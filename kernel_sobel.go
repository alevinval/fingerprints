package main

import "image"

var (
	SobelDx = &sobelKernel{size: 3, mat: [3][3]int{{-1, 0, 1}, {-2, 0, 2}, {-1, 0, 1}}}
	SobelDy = &sobelKernel{size: 3, mat: [3][3]int{{-1, -2, -1}, {0, 0, 0}, {1, 2, 1}}}
)

type sobelKernel struct {
	size int
	mat  [3][3]int
}

func (k *sobelKernel) Offset() int {
	return (k.size - 1) / 2
}

func (k *sobelKernel) Apply(in *image.Gray, x, y int) float64 {
	sum := 0
	for i := -k.Offset(); i <= k.Offset(); i++ {
		for j := -k.Offset(); j <= k.Offset(); j++ {
			sum += k.mat[j+1][i+1] * int(in.GrayAt(x+i, y+j).Y)
		}
	}
	return float64(sum)
}
