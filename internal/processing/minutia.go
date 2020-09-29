package processing

import (
	"github.com/alevinval/fingerprints/internal/matrix"
	"github.com/alevinval/fingerprints/internal/types"
)

func ExtractMinutia(skeleton *matrix.M, filteredDirectional *matrix.M, segmented *matrix.M) types.MinutiaeList {
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

func matchMinutiaeType(in *matrix.M, x, y int) types.MinutiaeType {
	p0 := in.At(x-1, y-1) > 0
	p1 := in.At(x, y-1) > 0
	p2 := in.At(x+1, y-1) > 0
	p3 := in.At(x+1, y) > 0
	p4 := in.At(x+1, y+1) > 0
	p5 := in.At(x, y+1) > 0
	p6 := in.At(x-1, y+1) > 0
	p7 := in.At(x-1, y) > 0
	pc := in.At(x, y) > 0
	f := func(pc, p0, p1, p2, p3, p4, p5, p6, p7 bool) bool {
		return pc && p0 && p1 && p2 && p3 && p4 && p5 && p6 && p7
	}

	/*
		p0 p1 p2
		p7 pc p3
		p6 p5 p4
	*/
	isBifurcation := (
	// Diagonals
	f(pc, p0, p3, p5, !p1, !p2, !p4, !p6, !p7) ||
		f(pc, p4, p1, p7, !p0, !p2, !p3, !p5, !p6) ||
		f(pc, p2, p7, p5, !p0, !p1, !p3, !p4, !p6) ||
		f(pc, p6, p1, p3, !p0, !p2, !p4, !p5, !p7) ||
		// Verticals/Horizontals
		f(pc, p1, p6, p4, !p0, !p2, !p3, !p5, !p7) ||
		f(pc, p3, p0, p6, !p1, !p2, !p4, !p5, !p7) ||
		f(pc, p5, p0, p2, !p1, !p3, !p4, !p6, !p7) ||
		f(pc, p6, p2, p4, !p0, !p1, !p3, !p5, !p7) ||
		// Perpendiculars clock and counter-clock wise
		f(pc, p1, p3, p5, !p0, !p2, !p4, !p6, !p7) ||
		f(pc, p1, p7, p5, !p0, !p2, !p3, !p4, !p6) ||
		f(pc, p3, p7, p5, !p0, !p1, !p2, !p4, !p6) ||
		f(pc, p3, p7, p1, !p0, !p2, !p4, !p5, !p6) ||
		f(pc, p5, p7, p1, !p0, !p2, !p3, !p4, !p6) ||
		f(pc, p5, p3, p1, !p0, !p2, !p4, !p6, !p7) ||
		f(pc, p7, p1, p3, !p0, !p2, !p4, !p5, !p6) ||
		f(pc, p7, p5, p3, !p0, !p1, !p2, !p4, !p6))

	isTermination := (f(pc, p0, !p1, !p2, !p3, !p4, !p5, !p6, !p7) ||
		f(pc, p1, !p0, !p2, !p3, !p4, !p5, !p6, !p7) ||
		f(pc, p2, !p0, !p1, !p3, !p4, !p5, !p6, !p7) ||
		f(pc, p3, !p0, !p1, !p2, !p4, !p5, !p6, !p7) ||
		f(pc, p4, !p0, !p1, !p2, !p3, !p5, !p6, !p7) ||
		f(pc, p5, !p0, !p1, !p2, !p3, !p4, !p6, !p7) ||
		f(pc, p6, !p0, !p1, !p2, !p3, !p4, !p5, !p7) ||
		f(pc, p7, !p0, !p1, !p2, !p3, !p4, !p5, !p6))

	if isBifurcation {
		return types.Bifurcation
	} else if isTermination {
		return types.Termination
	} else {
		return types.Unknown
	}
}
