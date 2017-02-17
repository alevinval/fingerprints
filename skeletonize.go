package main

func Skeletonize(in, out *Matrix) {
	bounds := in.Bounds()

	changes := true
	runs := 0
	for changes && runs < 20 {
		changes = false
		runs++

		conditions := []Condition{new(ConditionLeftBorder), new(ConditionRightBorder)}
		toRemove := [][2]int{}
		for _, c := range conditions {
			for x := bounds.Min.X + 1; x < bounds.Max.X-1; x++ {
				for y := bounds.Min.Y + 1; y < bounds.Max.Y-1; y++ {
					if in.At(x, y) > 0 {
						continue
					}
					p0 := in.At(x+1, y) == 0
					p1 := in.At(x+1, y-1) == 0
					p2 := in.At(x, y-1) == 0
					p3 := in.At(x-1, y-1) == 0
					p4 := in.At(x-1, y) == 0
					p5 := in.At(x-1, y+1) == 0
					p6 := in.At(x, y+1) == 0
					p7 := in.At(x+1, y+1) == 0
					if c.Holds(p0, p1, p2, p3, p4, p5, p6, p7) {
						toRemove = append(toRemove, [2]int{x, y})
					}
				}
			}

			if len(toRemove) > 0 {
				changes = true
			}
			for idx := range toRemove {
				in.Set(toRemove[idx][0], toRemove[idx][1], 255)
			}
		}
	}
}

type Condition interface {
	Holds(p0, p1, p2, p3, p4, p5, p6, p7 bool) bool
}

type ConditionLeftBorder int

func (c *ConditionLeftBorder) Holds(p0, p1, p2, p3, p4, p5, p6, p7 bool) bool {
	return !p4 && p0 && (p1 || p2 || p6 || p7) && (p2 || !p3) && (!p5 || p6)
}

type ConditionRightBorder int

func (c *ConditionRightBorder) Holds(p0, p1, p2, p3, p4, p5, p6, p7 bool) bool {
	return !p0 && p4 && (p2 || p3 || p5 || p6) && (p6 || !p7) && (!p1 || p2)
}

type ConditionTopBorder int

func (c *ConditionTopBorder) Holds(p0, p1, p2, p3, p4, p5, p6, p7 bool) bool {
	return !p2 && p6 && (p0 || p4 || p5 || p7) && (p0 || !p1) && (!p3 || p4)
}

type ConditionBottomBorder int

func (c *ConditionBottomBorder) Holds(p0, p1, p2, p3, p4, p5, p6, p7 bool) bool {
	return !p6 && p2 && (p0 || p1 || p3 || p4) && (p4 || !p5) && (p0 || !p7)
}
