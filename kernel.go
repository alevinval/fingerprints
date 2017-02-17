package main

import (
	"image"
	"sync"
)

type Kernel interface {
	Apply(in *Matrix, x, y int) float64
	Offset() int
}

func Convolute(k Kernel, in, out *Matrix) {
	offset := k.Offset()
	bounds := in.Bounds()
	for x := bounds.Min.X + offset; x < bounds.Max.X-offset; x++ {
		for y := bounds.Min.Y + offset; y < bounds.Max.Y-offset; y++ {
			pixel := k.Apply(in, x, y)
			out.Set(x, y, pixel)
		}
	}
}

func ParallelConvolution(k Kernel, in, out *Matrix) *sync.WaitGroup {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	sub := generateSubImageBounds(in)

	go func() {
		wgs := make([]*sync.WaitGroup, 0)
		for subImage := range sub {
			w := DeferredConvolution(k, subImage, out)
			wgs = append(wgs, w)
		}
		for _, w := range wgs {
			w.Wait()
		}
		wg.Done()
	}()

	return wg
}

func DeferredConvolution(k Kernel, in, out *Matrix) *sync.WaitGroup {
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		Convolute(k, in, out)
		wg.Done()
	}()
	return wg
}

func generateSubImageBounds(in *Matrix) <-chan *Matrix {
	ch := make(chan *Matrix)
	bounds := in.Bounds()
	blockSize := 200
	go func() {
		for x := bounds.Min.X; x < bounds.Max.X; x += blockSize {
			xp := x + blockSize
			if xp > bounds.Max.X {
				xp = bounds.Max.X
			}
			for y := bounds.Min.Y; y < bounds.Max.Y; y += blockSize {
				yp := y + blockSize
				if yp > bounds.Max.Y {
					yp = bounds.Max.Y
				}
				ch <- in.SubImage(image.Rect(x-1, y-1, xp+1, yp+1))
			}
		}
		close(ch)
	}()
	return ch
}
