package main

import (
	"image"
	"image/color"
	"math"
	"sync"
)

type Kernel interface {
	Apply(in *image.Gray, x, y int) int
	Offset() int
}

func ApplyKernelAsync(in *image.Gray, out *image.Gray, kernel Kernel) *sync.WaitGroup {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		ApplyKernel(kernel, in, out)
		wg.Done()
	}()
	return wg
}

func ApplyKernel(k Kernel, in *image.Gray, out *image.Gray) {
	var min, max int
	min = math.MaxInt64

	offset := k.Offset()
	bounds := in.Bounds()
	dx, dy := bounds.Dx(), bounds.Dy()
	for x := offset; x <= dx-offset; x++ {
		for y := offset; y <= dy-offset; y++ {
			val := k.Apply(in, x, y)
			if val > max {
				max = val
			}
			if val < min {
				min = val
			}
		}
	}
	for x := offset; x <= dx-offset; x++ {
		for y := offset; y <= dy-offset; y++ {
			val := k.Apply(in, x, y)
			normVal := uint8(math.MaxUint8 * float64(val-min) / float64(max-min))
			out.SetGray(x, y, color.Gray{Y: normVal})
		}
	}
}
