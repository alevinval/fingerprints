package main

import (
	"image"
	"image/color"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"os"
	_ "time"

	"github.com/nfnt/resize"
	_ "github.com/nfnt/resize"
)

func loadImage(name string) *image.Gray {
	f, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	m, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	m = resize.Resize(400, 400, m, resize.Bilinear)

	bounds := m.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y
	gray := image.NewGray(bounds)
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			oldColor := m.At(x, y)
			grayColor := color.GrayModel.Convert(oldColor)
			gray.Set(x, y, grayColor)
		}
	}

	return gray
}

var posX, posY = 0, 0

func showImage(title string, in *Matrix) {
	f, err := os.Create("out/" + title + ".png")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img := in.ToGray()
	png.Encode(f, img)
}

func processImage(in *Matrix) {
	bounds := in.Bounds()
	normalized := NewMatrix(bounds)

	Normalize(in, normalized)
	showImage("Normalized", normalized)

	gx, gy := NewMatrix(bounds), NewMatrix(bounds)
	c1 := ParallelConvolution(SobelDx, normalized, gx)
	c2 := ParallelConvolution(SobelDy, normalized, gy)
	c1.Wait()
	c2.Wait()

	//Consistency matrix
	consistency, normConsistency := NewMatrix(bounds), NewMatrix(bounds)
	c1 = ParallelConvolution(NewSqrtKernel(gx, gy), in, consistency)
	c1.Wait()
	Normalize(consistency, normConsistency)
	showImage("Normalized Consistency", normConsistency)

	// Compute directional
	directional, normDirectional := NewMatrix(bounds), NewMatrix(bounds)
	c1 = ParallelConvolution(NewDirectionalKernel(gx, gy), directional, directional)
	c1.Wait()
	Normalize(directional, normDirectional)
	showImage("Directional", normDirectional)

	// Compute filtered directional
	filteredD, normFilteredD := NewMatrix(bounds), NewMatrix(bounds)
	Convolute(NewFilteredDirectional(gx, gy, 4), filteredD, filteredD)
	Normalize(filteredD, normFilteredD)
	showImage("Filtered Directional", normFilteredD)

	// Compute segmented image
	segmented, normSegmented := NewMatrix(bounds), NewMatrix(bounds)
	Convolute(NewVarianceKernel(filteredD, 8), normalized, segmented)
	Normalize(segmented, normSegmented)
	showImage("Filtered Directional Std Dev.", normSegmented)

	// Compute binarized segmented image
	binarizedSegmented := NewMatrix(bounds)
	Binarize(normSegmented, binarizedSegmented)
	showImage("Binarized Segmented", binarizedSegmented)

	// Binarize normalized image
	binarizedNorm := NewMatrix(bounds)
	Binarize(normalized, binarizedNorm)
	showImage("Binarized Normalized", binarizedNorm)

	BinarizeEnhancement(binarizedNorm)
	showImage("Binarized Enhanced", binarizedNorm)

	// Skeletonize
	Skeletonize(binarizedNorm)
	showImage("Skeletonized", binarizedNorm)
}

func main() {
	original := loadImage("corpus/nist2.jpg")
	img := NewMatrixFromGray(original)
	processImage(img)
}
