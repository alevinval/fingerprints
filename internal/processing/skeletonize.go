package processing

import "github.com/alevinval/fingerprints/internal/matrix"

func Skeletonize(in *matrix.M) {
	bounds := in.Bounds()

	changes := true
	for changes {
		changes = false
		for _, c := range conditions {
			toRemove := [][2]int{}
			for y := bounds.Min.Y + 1; y < bounds.Max.Y-1; y++ {
				for x := bounds.Min.X + 1; x < bounds.Max.X-1; x++ {
					if in.At(x, y) != black {
						continue
					}
					p0 := in.At(x+1, y) == black
					p1 := in.At(x+1, y+1) == black
					p2 := in.At(x, y+1) == black
					p3 := in.At(x-1, y+1) == black
					p4 := in.At(x-1, y) == black
					p5 := in.At(x-1, y-1) == black
					p6 := in.At(x, y-1) == black
					p7 := in.At(x+1, y-1) == black

					if c.Holds(p0, p1, p2, p3, p4, p5, p6, p7) {
						toRemove = append(toRemove, [2]int{x, y})
					}
				}
			}
			changes = len(toRemove) > 0
			for idx := range toRemove {
				in.Set(toRemove[idx][0], toRemove[idx][1], white)
			}
		}
	}
}

var conditions = []Condition{
	new(cLeftBorder), new(cRightBorder),
	new(cTopBorder), new(cBottomBorder),
}

type (
	cLeftBorder   struct{}
	cRightBorder  struct{}
	cTopBorder    struct{}
	cBottomBorder struct{}

	Condition interface {
		Holds(p0, p1, p2, p3, p4, p5, p6, p7 bool) bool
	}
)

func (c cLeftBorder) Holds(p0, p1, p2, p3, p4, p5, p6, p7 bool) bool {
	return !p4 && p0 && (p1 || p2 || p6 || p7) && (p2 || !p3) && (!p5 || p6)
}

func (c cRightBorder) Holds(p0, p1, p2, p3, p4, p5, p6, p7 bool) bool {
	return !p0 && p4 && (p2 || p3 || p5 || p6) && (p6 || !p7) && (!p1 || p2)
}

func (c cTopBorder) Holds(p0, p1, p2, p3, p4, p5, p6, p7 bool) bool {
	return !p2 && p6 && (p0 || p4 || p5 || p7) && (p0 || !p1) && (!p3 || p4)
}

func (c cBottomBorder) Holds(p0, p1, p2, p3, p4, p5, p6, p7 bool) bool {
	return !p6 && p2 && (p0 || p1 || p3 || p4) && (p4 || !p5) && (p0 || !p7)
}
