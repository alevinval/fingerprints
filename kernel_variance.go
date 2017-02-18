package main

import "math"

type varianceKernel struct {
	phy    *Matrix
	offset int
}

func NewVarianceKernel(directional *Matrix, offset int) *varianceKernel {
	return &varianceKernel{phy: directional, offset: offset}
}

func (k *varianceKernel) Offset() int {
	return k.offset
}

func (k *varianceKernel) Apply(in *Matrix, x, y int) float64 {
	var pos int

	sigSize := float64(k.Offset()*2 + 1)
	signature := make([]float64, int(sigSize))
	angle := k.phy.At(x, y) - math.Pi/2
	for i := x - k.Offset(); i <= x+k.Offset(); i++ {
		for j := y - k.Offset(); j <= y+k.Offset(); j++ {
			xp := int(float64(i-x)*math.Cos(angle)-float64(j-y)*math.Sin(angle)) + x
			yp := int(float64(i-x)*math.Sin(angle)+float64(j-y)*math.Cos(angle)) + y
			if xp >= 0 && xp < in.bounds.Dx() && yp >= 0 && yp < in.bounds.Dy() {
				signature[pos] += in.At(xp, yp)
			} else {
				signature[pos] += in.At(x, y)
			}
		}
		pos++
	}

	sum := 0.0
	for _, sig := range signature {
		sum += sig
	}
	mean := sum / sigSize

	variance := 0.0
	for _, sig := range signature {
		d := sig - mean
		variance += d * d
	}
	variance /= sigSize
	return variance
}
