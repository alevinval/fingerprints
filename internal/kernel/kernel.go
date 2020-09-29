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

func (base *Base) ConvoluteParallelized(in, out *matrix.M) {
	subImages := generateSubImages(in, base.kernel.Offset())

	wg := &sync.WaitGroup{}
	wg.Add(len(subImages))
	for _, image := range subImages {
		go base.convoluteWithWG(wg, image, out)
	}
	wg.Wait()
}

func (base *Base) Convolute(in, out *matrix.M) {
	offset := base.kernel.Offset()
	bounds := in.Bounds()
	for y := bounds.Min.Y + offset; y < bounds.Max.Y-offset; y++ {
		for x := bounds.Min.X + offset; x < bounds.Max.X-offset; x++ {
			pixel := base.kernel.Apply(in, x, y)
			out.Set(x, y, pixel)
		}
	}
}

func (base *Base) convoluteWithWG(wg *sync.WaitGroup, in, out *matrix.M) {
	base.Convolute(in, out)
	wg.Done()
}

func generateSubImages(in *matrix.M, offset int) []*matrix.M {
	images := []*matrix.M{}
	bounds := in.Bounds()
	blockSize := bounds.Max.X / 2
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
			image := in.SubImage(image.Rect(xi, yi, xp, yp))
			images = append(images, image)
		}
	}
	return images
}
