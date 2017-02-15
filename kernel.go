package main

import "sync"

type Kernel interface {
	Apply(in *Matrix, x, y int) float64
	Offset() int
}

func DeferredConvolution(k Kernel, in *Matrix, out *Matrix) *sync.WaitGroup {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		Convolute(k, in, out)
		wg.Done()
	}()
	return wg
}

func Convolute(k Kernel, in *Matrix, out *Matrix) {
	offset := k.Offset()
	bounds := in.Bounds()
	dx, dy := bounds.Dx(), bounds.Dy()
	for x := offset; x < dx-offset; x++ {
		for y := offset; y < dy-offset; y++ {
			pixel := k.Apply(in, x, y)
			out.Set(x, y, pixel)
		}
	}
}
