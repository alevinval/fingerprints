package processing

import (
	"image"
	"log"
	"math"
	"sync"

	"github.com/alevinval/fingerprints/internal/helpers"
	"github.com/alevinval/fingerprints/internal/matrix"
	"github.com/alevinval/fingerprints/internal/types"
)

const (
	black = 0
	white = 255
)

// BinarizeSegmented runs binarization with an optimized threshold that
// ensures the segmented area is as big an continuous as possible
func BinarizeSegmented(in, out *matrix.M, meta types.Metadata) {
	binarize(in, out, math.Sqrt(meta.MeanValue))
}

// BinarizeSkeleton runs binarization with an optimized threshold that
// does not damage the skeleton itself.
func BinarizeSkeleton(in, out *matrix.M, meta types.Metadata) {
	binarize(in, out, meta.MeanValue/(math.Pi/2))
}

func binarize(in, out *matrix.M, threshold float64) {
	helpers.RunInParallel(in, 0, func(wg *sync.WaitGroup, bounds image.Rectangle) {
		doBinarize(in, out, bounds, threshold)
		wg.Done()
	})
}

func doBinarize(in *matrix.M, out *matrix.M, bounds image.Rectangle, threshold float64) {
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			if in.At(x, y) < threshold {
				out.Set(x, y, black)
			} else {
				out.Set(x, y, white)
			}
		}
	}
}

func BinarizeEnhancement(in *matrix.M) *matrix.M {
	bounds := in.Bounds()
	p := matrix.NewFromGray(in.ToGray())

	region := 1
	for y := bounds.Min.Y + 1; y < bounds.Max.Y-1; y++ {
		for x := bounds.Min.X + 1; x < bounds.Max.X-1; x++ {
			if p.At(x, y) == black || p.At(x, y) == white {
				fillRegion(p, region, x, y, 0)
				region++
			}
			if region == white {
				region++
			}
		}
	}
	log.Printf("regions found: %d", region)
	histogram := make([]int, region)
	for y := bounds.Min.Y + 1; y < bounds.Max.Y-1; y++ {
		for x := bounds.Min.X + 1; x < bounds.Max.X-1; x++ {
			histogram[int(p.At(x, y))] += 1
		}
	}

	sum := 0.0
	for _, area := range histogram {
		sum += float64(area)
	}

	mean := sum / float64(region)

	erasedRegions := 0
	for region, area := range histogram {
		if float64(area) < math.Sqrt(mean) {
			eraseRegion(p, in, region)
			erasedRegions++
		}
	}
	log.Printf("erased regions: %d", erasedRegions)
	return p
}

func fillRegion(p *matrix.M, region, x, y, max int) {
	if x == p.Bounds().Min.X+1 || x == p.Bounds().Max.X-1 {
		return
	}
	if y == p.Bounds().Min.Y+1 || y == p.Bounds().Max.Y-1 {
		return
	}
	bw := p.At(x, y)
	p.Set(x, y, float64(region))
	if p.At(x-1, y) == bw {
		fillRegion(p, region, x-1, y, max)
	}
	if p.At(x, y-1) == bw {
		fillRegion(p, region, x, y-1, max)
	}
	if p.At(x+1, y) == bw {
		fillRegion(p, region, x+1, y, max)
	}
	if p.At(x, y+1) == bw {
		fillRegion(p, region, x, y+1, max)
	}
	if p.At(x-1, y-1) == bw {
		fillRegion(p, region, x-1, y-1, max)
	}
	if p.At(x+1, y-1) == bw {
		fillRegion(p, region, x+1, y-1, max)
	}
	if p.At(x+1, y+1) == bw {
		fillRegion(p, region, x+1, y+1, max)
	}
	if p.At(x-1, y+1) == bw {
		fillRegion(p, region, x-1, y+1, max)
	}
}

func eraseRegion(p, in *matrix.M, region int) {
	bounds := p.Bounds()
	for y := bounds.Min.Y + 1; y < bounds.Max.Y-1; y++ {
		for x := bounds.Min.X + 1; x < bounds.Max.X-1; x++ {
			value := p.At(x, y)
			if int(value) != region {
				continue
			} else if value == white {
				in.Set(x, y, black)
			} else {
				in.Set(x, y, white)
			}
		}
	}
}
