package extraction

import (
	"github.com/alevinval/fingerprints/internal/matrix"
	"github.com/alevinval/fingerprints/internal/types"
)

// Minutia retrieves fingerprint features from a skeletonized image. Each
// feature angle is obtained from the filtered directional image. Features
// outside the fingerprint itself are removed by checking against the
// segmented image, that tells us what is fingerprint from background.
func Minutia(skeleton *matrix.M, filteredDirectional *matrix.M, segmented *matrix.M) types.MinutiaeList {
	minutiaes := types.MinutiaeList{}
	bounds := skeleton.Bounds()
	for y := bounds.Min.Y + 1; y < bounds.Max.Y-1; y++ {
		for x := bounds.Min.X + 1; x < bounds.Max.X-1; x++ {
			if segmented.At(x, y) == 0 {
				continue
			}
			minutiaeType := matchMinutiaeType(skeleton, x, y)
			if minutiaeType != types.Unknown {
				minutiae := types.Minutiae{
					X:     x,
					Y:     y,
					Angle: filteredDirectional.At(x, y),
					Type:  minutiaeType,
				}
				minutiaes = append(minutiaes, minutiae)
			}
		}
	}
	return minutiaes
}

func matchMinutiaeType(in *matrix.M, i, j int) types.MinutiaeType {
	p0 := in.At(i-1, j-1) > 0
	p1 := in.At(i, j-1) > 0
	p2 := in.At(i+1, j-1) > 0
	p3 := in.At(i+1, j) > 0
	p4 := in.At(i+1, j+1) > 0
	p5 := in.At(i, j+1) > 0
	p6 := in.At(i-1, j+1) > 0
	p7 := in.At(i-1, j) > 0
	pc := in.At(i, j) > 0

	and := func(f0, f1, f2, f7, fc, f3, f6, f5, f4 bool) bool {
		return (pc == fc) && (p0 == f0) && (p1 == f1) && (p2 == f2) && (p3 == f3) && (p4 == f4) && (p5 == f5) && (p6 == f6) && (p7 == f7)
	}

	isPore := and(o, x, o,
		x, o, x,
		o, x, o)

	if isPore {
		return types.Pore
	}

	isBifurcation := (
	// Diagonals
	and(x, o, x,
		o, x, o,
		o, o, x) ||
		and(x, o, x,
			o, x, o,
			o, x, o) ||
		and(x, o, x,
			o, x, o,
			x, o, o) ||
		and(x, o, o,
			o, x, x,
			x, o, o) ||
		and(x, o, o,
			o, x, o,
			x, o, x) ||
		and(o, x, o,
			o, x, o,
			x, o, x) ||
		and(o, o, x,
			o, x, o,
			x, o, x) ||
		and(o, o, x,
			x, x, o,
			o, o, x) ||
		and(x, o, x,
			o, x, o,
			o, o, x) ||
		// Orthogonals
		and(o, o, o,
			x, x, x,
			o, x, o) ||
		and(o, x, o,
			o, x, x,
			o, x, o) ||
		and(o, x, o,
			x, x, x,
			o, o, o) ||
		and(o, x, o,
			x, x, o,
			o, x, o) ||
		and(x, o, o,
			o, x, x,
			o, x, o) ||
		and(o, x, o,
			x, x, o,
			o, o, x) ||
		and(o, o, x,
			x, x, o,
			o, x, o) ||
		and(o, x, o,
			o, x, x,
			x, o, o))

	if isBifurcation {
		return types.Bifurcation
	}

	isTermination := (
	// Terminations
	and(x, o, o,
		o, x, o,
		o, o, o) ||
		and(o, x, o,
			o, x, o,
			o, o, o) ||
		and(o, o, x,
			o, x, o,
			o, o, o) ||
		and(o, o, o,
			o, x, x,
			o, o, o) ||
		and(o, o, o,
			o, x, o,
			o, o, x) ||
		and(o, o, o,
			o, x, o,
			o, x, o) ||
		and(o, o, o,
			o, x, o,
			x, o, o) ||
		and(o, o, o,
			x, x, o,
			o, o, o))

	if isTermination {
		return types.Termination
	}

	return types.Unknown
}

const (
	x = true
	o = false
)
