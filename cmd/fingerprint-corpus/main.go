package main

import (
	"image"
	"image/color"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"os"
	_ "time"

	"github.com/alevinval/fingerprints/internal/kernel"
	"github.com/alevinval/fingerprints/internal/matrix"
	"github.com/alevinval/fingerprints/internal/processing"
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

func showImage(title string, in *matrix.M) {
	f, err := os.Create("out2/" + title + ".png")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img := in.ToGray()
	png.Encode(f, img)
}

func processImage(in *matrix.M) {
	bounds := in.Bounds()
	normalized := matrix.New(bounds)

	processing.Normalize(in, normalized)
	showImage("Normalized", normalized)

	gx, gy := matrix.New(bounds), matrix.New(bounds)
	c1 := kernel.SobelDx.ParallelConvolution(normalized, gx)
	c2 := kernel.SobelDy.ParallelConvolution(normalized, gy)
	c1.Wait()
	c2.Wait()

	//Consistency matrix
	consistency, normConsistency := matrix.New(bounds), matrix.New(bounds)
	c1 = kernel.NewSqrt(gx, gy).ParallelConvolution(in, consistency)
	c1.Wait()
	processing.Normalize(consistency, normConsistency)
	showImage("Normalized Consistency", normConsistency)

	// Compute directional
	directional, normDirectional := matrix.New(bounds), matrix.New(bounds)
	c1 = kernel.NewDirectional(gx, gy).ParallelConvolution(directional, directional)
	c1.Wait()
	processing.Normalize(directional, normDirectional)
	showImage("Directional", normDirectional)

	// Compute filtered directional
	filteredD, normFilteredD := matrix.New(bounds), matrix.New(bounds)
	c1 = kernel.FilteredDirectional(gx, gy, 4).ParallelConvolution(filteredD, filteredD)
	c1.Wait()
	processing.Normalize(filteredD, normFilteredD)
	showImage("Filtered Directional", normFilteredD)

	// Compute segmented image
	segmented, normSegmented := matrix.New(bounds), matrix.New(bounds)
	kernel.Variance(filteredD).Convolution(normalized, segmented)
	processing.Normalize(segmented, normSegmented)
	showImage("Filtered Directional Std Dev.", normSegmented)

	// Compute binarized segmented image
	binarizedSegmented := matrix.New(bounds)
	processing.Binarize(normSegmented, binarizedSegmented)
	showImage("Binarized Segmented", binarizedSegmented)

	// Binarize normalized image
	binarizedNorm := matrix.New(bounds)
	processing.Binarize(normalized, binarizedNorm)
	showImage("Binarized Normalized", binarizedNorm)

	processing.BinarizeEnhancement(binarizedNorm)
	showImage("Binarized Enhanced", binarizedNorm)

	// Skeletonize
	processing.Skeletonize(binarizedNorm)
	showImage("Skeletonized", binarizedNorm)
}

func main() {
	original := loadImage("corpus/nist2.jpg")
	img := matrix.NewFromGray(original)
	processImage(img)
}
