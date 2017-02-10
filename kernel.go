package main

import (
	"image"
	"image/color"
	"math"
	"sync"
)

var (
	SobelDx = &kernel3x3{mat: [3][3]int{{-1, 0, 1}, {-2, 0, 2}, {-1, 0, 1}}}
	SobelDy = &kernel3x3{mat: [3][3]int{{-1, -2, -1}, {0, 0, 0}, {1, 2, 1}}}
)

type Kernel interface {
	Apply(in *image.Gray, x, y int) int
}

type kernel3x3 struct {
	mat [3][3]int
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

func ApplyKernel(kernel Kernel, in *image.Gray, out *image.Gray) {
	var min, max int
	min = math.MaxInt64

	for x := 1; x <= in.Bounds().Dx()-1; x++ {
		for y := 1; y <= in.Bounds().Dy()-1; y++ {
			val := kernel.Apply(in, x, y)
			if val > max {
				max = val
			}
			if val < min {
				min = val
			}
		}
	}
	for x := 1; x <= in.Bounds().Dx()-1; x++ {
		for y := 1; y <= in.Bounds().Dy()-1; y++ {
			val := kernel.Apply(in, x, y)
			normVal := uint8(math.MaxUint8 * float64(val-min) / float64(max-min))
			out.SetGray(x, y, color.Gray{Y: normVal})
		}
	}
}

func (sk *kernel3x3) Apply(in *image.Gray, x, y int) int {
	sum := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			sum += sk.mat[j+1][i+1] * int(in.GrayAt(x+i, y+j).Y)
		}
	}
	return sum
}
