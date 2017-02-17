package main

import "math"

type stdDevKernel struct {
	phy *Matrix
	offset int
}

func NewStdDevKernel(directional *Matrix, offset int) *stdDevKernel {
	return &stdDevKernel{phy: directional, offset: offset}
}

func (k *stdDevKernel) Offset() int {
	return k.offset
}

func (k *stdDevKernel) Apply(in *Matrix, x, y int) float64 {
	var pos int

	signature_size := float64(k.Offset() * 2 + 1)
	signature := make([]float64, int(signature_size))
	angle := k.phy.At(x, y) - math.Pi/2
	for i := x-k.Offset(); i <= x+k.Offset(); i++ {
		for j := y-k.Offset(); j <= y+k.Offset(); j++ {
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
	mean := sum / signature_size
	variance := 0.0
	for _, sig := range signature {
		d := sig - mean
		variance += d * d
	}
	variance /= signature_size
	return math.Sqrt(variance)
}
