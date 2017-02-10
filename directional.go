package main

import (
	"image"
	"image/color"
	"math"
)

func ComputeDirectional(in *image.Gray) *image.Gray {
	rect := in.Bounds()
	out := image.NewGray(rect)
	Gx := image.NewGray(rect)
	Gy := image.NewGray(rect)

	w1 := ApplyKernelAsync(in, Gx, SobelDx)
	w2 := ApplyKernelAsync(in, Gy, SobelDy)
	w1.Wait()
	w2.Wait()

	for x := 1; x <= rect.Dx()-1; x++ {
		for y := 1; y <= rect.Dy()-1; y++ {
			dx := Gx.GrayAt(x, y).Y
			dy := Gy.GrayAt(x, y).Y
			ang := math.Atan2(float64(dy), float64(dx))
			val := uint8((ang + math.Pi/2) / math.Pi * 255)
			out.SetGray(x, y, color.Gray{Y: val})
		}
	}
	return out
}
