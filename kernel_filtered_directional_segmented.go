package main

import "math"

type fdKernelSegmented struct {
	phy *Matrix
}

func NewFilteredDirectionalSegmented(directional *Matrix) *fdKernelSegmented {
	return &fdKernelSegmented{phy: directional}
}

func (k *fdKernelSegmented) Offset() int {
	return 8
}

func (k *fdKernelSegmented) Apply(in *Matrix, x, y int) float64 {
	var pos int
	bounds := in.Bounds()

	signature := make([]float64, k.Offset()*2)
	angle := k.phy.At(x, y) - math.Pi/2
	for i := -k.Offset(); i < k.Offset(); i++ {
		for j := -k.Offset(); j < k.Offset(); j++ {
			xp := int(math.Ceil(float64(i-x)*math.Cos(angle)-float64(j-y)*math.Sin(angle)+float64(x))) + x
			yp := int(math.Ceil(float64(i-x)*math.Sin(angle)+float64(j-y)*math.Cos(angle)+float64(y))) + y
			if xp < 0 || xp >= bounds.Dx() {
				//signature[pos] += in.At(x, y)
				continue
			}
			if yp < 0 || yp >= bounds.Dy() {
				//signature[pos] += in.At(x, y)
				continue
			}
			signature[pos] += in.At(xp, yp)
		}
		pos++
	}

	sum := 0.0
	for _, sig := range signature {
		sum += sig
	}
	mean := sum / float64(k.Offset()*2)
	variance := 0.0
	for _, sig := range signature {
		d := sig - mean
		variance += d * d
	}
	variance /= float64(k.Offset() * 2)
	return variance
	//std := math.Sqrt(variance)
	//return std
	//if variance > 25000 {
	//	return 255
	//}
	//return 0
}
