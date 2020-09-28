package main

import (
	"image"
	"image/color"
	_ "image/jpeg"
	"image/png"
	"log"
	"os"
	"path"

	"github.com/alevinval/fingerprints/internal/kernel"
	"github.com/alevinval/fingerprints/internal/matrix"
	"github.com/alevinval/fingerprints/internal/processing"
	"github.com/nfnt/resize"
)

var outFolder = "out"

func resizeImage(img image.Image) image.Image {
	maxDimension := 300
	dx := img.Bounds().Dx()
	dy := img.Bounds().Dy()

	xp := 0
	yp := 0
	if dx > dy {
		xp = maxDimension
		yp = int(float64(dy) / (float64(dx) / float64(maxDimension)))
	} else if dy > dx {
		yp = maxDimension
		xp = int(float64(dx) / (float64(dy) / float64(maxDimension)))
	} else {
		xp, yp = maxDimension, maxDimension
	}
	log.Printf("resizing image from (%d,%d) to (%d,%d)", dx, dy, xp, yp)
	return resize.Resize(uint(xp), uint(yp), img, resize.Bilinear)
}

func loadImage(name string) *image.Gray {
	f, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	resizedImg := resizeImage(img)

	bounds := resizedImg.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y
	gray := image.NewGray(bounds)
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			oldColor := resizedImg.At(x, y)
			grayColor := color.GrayModel.Convert(oldColor)
			gray.Set(x, y, grayColor)
		}
	}

	return gray
}

var posX, posY = 0, 0

func showImage(title string, in *matrix.M) {
	f, err := os.Create(path.Join(outFolder, title+".png"))
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
	kernel.SobelDx.ParallelConvolution(normalized, gx)
	kernel.SobelDy.ParallelConvolution(normalized, gy)

	// Compute filtered directional
	filteredD, normFilteredD := matrix.New(bounds), matrix.New(bounds)
	kernel.FilteredDirectional(gx, gy, 4).ParallelConvolution(filteredD, filteredD)
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
	processing.BinarizeEnhancement(binarizedSegmented)
	showImage("Binarized Segmented", binarizedSegmented)

	// Binarize normalized image
	skeletonized := matrix.New(bounds)
	processing.Binarize(normalized, skeletonized)
	processing.BinarizeEnhancement(skeletonized)
	processing.Skeletonize(skeletonized)
	showImage("Skeletonized", skeletonized)

	minutiaes := processing.ExtractMinutiae(skeletonized, filteredD, binarizedSegmented)

	out := matrix.New(bounds)
	for _, minutiae := range minutiaes {
		log.Printf("Found minutiae at %d, %d", minutiae.X, minutiae.Y)
		log.Printf("Type=%v, Angle=%f", minutiae.Type, minutiae.Angle)
		out.Set(minutiae.X, minutiae.Y, 255.0)
	}

	showImage("Minutiaes", out)
}

func main() {
	original := loadImage("corpus/nist2.jpg")
	img := matrix.NewFromGray(original)
	processImage(img)
}
