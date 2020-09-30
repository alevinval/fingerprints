package processing

import (
	"log"
	"math"

	"github.com/alevinval/fingerprints/internal/matrix"
	"github.com/alevinval/fingerprints/internal/types"
)

const BLACK = 0
const WHITE = 255

func Binarize(in, out *matrix.M, meta types.Metadata) {
	bounds := in.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			pixel := in.At(x, y)
			if float64(pixel) < meta.MeanValue/(math.Pi/2) {
				out.Set(x, y, BLACK)
			} else {
				out.Set(x, y, WHITE)
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
			if p.At(x, y) == BLACK || p.At(x, y) == WHITE {
				fillRegion(p, region, x, y, 0)
				region++
			}
			if region == WHITE {
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
			} else if value == WHITE {
				in.Set(x, y, BLACK)
			} else {
				in.Set(x, y, WHITE)
			}
		}
	}
}
