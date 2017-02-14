package main

import "image"

var (
	normalize = &normalizeKernel{}
)

type normalizeKernel struct{}

func (k *normalizeKernel) Offset() int {
	return 0
}

func (k *normalizeKernel) Apply(in *image.Gray, x, y int) int {
	return int(in.GrayAt(x, y).Y)
}
