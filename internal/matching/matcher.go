package matching

import "math"

func Match(l1, l2 MinutiaeList) MinutiaeList {
	matches := MinutiaeList{}
	matched := map[Minutiae]struct{}{}

	for _, minutiae := range l1 {
		for _, candidate := range l2 {
			if _, ok := matched[candidate]; ok {
				continue
			}
			if minutiae.Type != candidate.Type {
				continue
			}
			if minutiae.Angle-candidate.Angle > 0.01 {
				continue
			}
			if distance(minutiae, candidate) > 5 {
				continue
			}
			matched[candidate] = struct{}{}
			matches = append(matches, minutiae)
		}
	}

	return matches
}

func distance(a, b Minutiae) float64 {
	dx := float64(b.X - a.X)
	dy := float64(b.Y - a.Y)
	return math.Sqrt(dx*dx + dy*dy)
}
