package main

import (
	"image"
	"image/color"
	"math"
)

func ComputeDirectional(in *image.Gray) {
	rect := in.Bounds()
	Gx := image.NewGray(rect)
	Gy := image.NewGray(rect)

	w1 := ApplyKernelAsync(in, Gx, SobelDx)
	w2 := ApplyKernelAsync(in, Gy, SobelDy)
	w1.Wait()
	w2.Wait()

	for x := SobelDy.size; x <= rect.Dx()-1; x++ {
		for y := 1; y <= rect.Dy()-1; y++ {
			dx := Gx.GrayAt(x, y).Y
			dy := Gy.GrayAt(x, y).Y

			ang := math.Atan2(float64(dy), float64(dx))
			val := uint8((ang + math.Pi/2) / math.Pi * 255)
			in.SetGray(x, y, color.Gray{Y: val})
		}
	}

	ComputeFilteredDirectional(Gx, Gy, in)
}

func ComputeFilteredDirectional(Gx *image.Gray, Gy *image.Gray, out *image.Gray) {
	rect := Gx.Bounds()

	var min, max float64
	min = math.MaxFloat64
	for x := Sum9x9.size / 2; x <= rect.Dx()-Sum9x9.size/2; x++ {
		for y := Sum9x9.size / 2; y <= rect.Dy()-Sum9x9.size/2; y++ {
			dx := Sum9x9.Apply(Gx, x, y)
			dy := Sum9x9.Apply(Gy, x, y)
			val := math.Pi/2 + 0.5*math.Atan2(float64(2*dx*dx), float64(dx*dx-dy*dy))
			if val < min {
				min = val
			}
			if val > max {
				max = val
			}

		}
	}
	for x := Sum9x9.size / 2; x <= rect.Dx()-Sum9x9.size/2; x++ {
		for y := Sum9x9.size / 2; y <= rect.Dy()-Sum9x9.size/2; y++ {
			dx := Sum9x9.Apply(Gx, x, y)
			dy := Sum9x9.Apply(Gy, x, y)
			val := math.Pi/2 + 0.5*math.Atan2(float64(2*dx*dx), float64(dx*dx-dy*dy))
			normVal := uint8(255 * (val - min) / (max - min))
			out.SetGray(x, y, color.Gray{Y: normVal})
		}
	}
}
