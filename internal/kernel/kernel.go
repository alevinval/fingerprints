package kernel

import (
	"image"
	"sync"

	"github.com/alevinval/fingerprints/internal/matrix"
)

type Kernel interface {
	Apply(in *matrix.M, x, y int) float64
	Offset() int
}

type Base struct {
	kernel Kernel
}

func (base *Base) ParallelConvolution(in, out *matrix.M) {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	sub := generateSubImageBounds(in, base.kernel.Offset())

	go func() {
		wgs := make([]*sync.WaitGroup, 0)
		for subImage := range sub {
			w := base.deferredConvolution(subImage, out)
			wgs = append(wgs, w)
		}
		for _, w := range wgs {
			w.Wait()
		}
		wg.Done()
	}()
	wg.Wait()
}

func (base *Base) Convolution(in, out *matrix.M) {
	offset := base.kernel.Offset()
	bounds := in.Bounds()
	for y := bounds.Min.Y + offset; y < bounds.Max.Y-offset; y++ {
		for x := bounds.Min.X + offset; x < bounds.Max.X-offset; x++ {
			pixel := base.kernel.Apply(in, x, y)
			out.Set(x, y, pixel)
		}
	}
}

func (base *Base) deferredConvolution(in, out *matrix.M) *sync.WaitGroup {
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		base.Convolution(in, out)
		wg.Done()
	}()
	return wg
}

func generateSubImageBounds(in *matrix.M, offset int) <-chan *matrix.M {
	ch := make(chan *matrix.M)
	bounds := in.Bounds()
	blockSize := bounds.Max.X / 2
	go func() {
		for x := bounds.Min.X; x < bounds.Max.X; x += blockSize {
			xi := x - offset
			if xi < bounds.Min.X {
				xi = bounds.Min.X
			}
			xp := x + blockSize + offset
			if xp > bounds.Max.X {
				xp = bounds.Max.X
			}
			for y := bounds.Min.Y; y < bounds.Max.Y; y += blockSize {
				yi := y - offset
				if yi < bounds.Min.Y {
					yi = bounds.Min.Y
				}
				yp := y + blockSize + offset
				if yp > bounds.Max.Y {
					yp = bounds.Max.Y
				}
				ch <- in.SubImage(image.Rect(xi, yi, xp, yp))
			}
		}
		close(ch)
	}()
	return ch
}
