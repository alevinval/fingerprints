package kernel

import (
	"image"
	"sync"

	"github.com/alevinval/fingerprints/internal/helpers"
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
	subBounds := helpers.GenerateSubBounds(in, base.kernel.Offset())

	wg := new(sync.WaitGroup)
	wg.Add(len(subBounds))
	for _, bounds := range subBounds {
		go base.convoluteWithWG(wg, in, out, bounds)
	}
	wg.Wait()
}

func (base *Base) Convolute(in, out *matrix.M) {
	base.convoluteWithBounds(in, out, in.Bounds())
}

func (base *Base) convoluteWithBounds(in, out *matrix.M, bounds image.Rectangle) {
	offset := base.kernel.Offset()
	for y := bounds.Min.Y + offset; y < bounds.Max.Y-offset; y++ {
		for x := bounds.Min.X + offset; x < bounds.Max.X-offset; x++ {
			pixel := base.kernel.Apply(in, x, y)
			out.Set(x, y, pixel)
		}
	}
}

func (base *Base) convoluteWithWG(wg *sync.WaitGroup, in, out *matrix.M, subBounds image.Rectangle) {
	base.convoluteWithBounds(in, out, subBounds)
	wg.Done()
}
