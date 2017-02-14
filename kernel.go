package main

import (
	"image"
	"image/color"
	"math"
	"sync"
)

type Kernel interface {
	Apply(in *image.Gray, x, y int) float64
	Offset() int
}

func DeferredConvolution(k Kernel, in *image.Gray, out *image.Gray) *sync.WaitGroup {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		Convolute(k, in, out)
		wg.Done()
	}()
	return wg
}

func Convolute(k Kernel, in *image.Gray, out *image.Gray) {
	var min, max float64
	min = math.MaxFloat64

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
			normVal := uint8(math.MaxUint8 * (val - min) / (max - min))
			out.SetGray(x, y, color.Gray{Y: normVal})
		}
	}
}
