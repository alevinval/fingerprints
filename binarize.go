package main

func Binarize(in, out *Matrix) {
	var sum float64

	bounds := in.Bounds()
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			sum += in.At(x, y)
		}
	}

	mean := sum / float64(bounds.Dx()*bounds.Dy())
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			pixel := in.At(x, y)
			if pixel < mean/3 {
				out.Set(x, y, 0)
			} else {
				out.Set(x, y, 255)
			}
		}
	}
}
