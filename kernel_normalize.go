package main

import "image"

var (
	normalize = &normalizeKernel{}
)

type normalizeKernel struct{}

func (k *normalizeKernel) Offset() int {
	return 0
}

func (k *normalizeKernel) Apply(in *image.Gray, x, y int) float64 {
	return float64(in.GrayAt(x, y).Y)
}
