package main

import "image"

var (
	Sum9x9 = &sumKernel{size: 9}
)

type sumKernel struct {
	size int
}

func (k *sumKernel) Offset() int {
	if k.size%2 == 0 {
		return k.size / 2
	}
	return (k.size - 1) / 2
}

func (k *sumKernel) Apply(in *image.Gray, x, y int) float64 {
	sum := 0
	for i := -k.Offset(); i <= k.Offset(); i++ {
		for j := -k.Offset(); j <= k.Offset(); j++ {
			sum += int(in.GrayAt(x+i, y+j).Y)
		}
	}
	return float64(sum)
}
