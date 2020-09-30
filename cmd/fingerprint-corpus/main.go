package main

import (
	"image"
	"image/png"
	"log"
	"os"
	"path"

	"github.com/alevinval/fingerprints/internal/debug"
	"github.com/alevinval/fingerprints/internal/extraction"
	"github.com/alevinval/fingerprints/internal/helpers"
	"github.com/alevinval/fingerprints/internal/kernel"
	"github.com/alevinval/fingerprints/internal/matrix"
	"github.com/alevinval/fingerprints/internal/processing"
)

var outFolder = "out"

func showImage(title string, img image.Image) {
	f, err := os.Create(path.Join(outFolder, title+".png"))
	if err != nil {
		panic(err)
	}
	defer f.Close()

	png.Encode(f, img)
}

func showMatrix(title string, m *matrix.M) {
	showImage(title, m.ToGray())
}

func processImage(img image.Image, in *matrix.M) {
	bounds := in.Bounds()
	normalized := matrix.New(bounds)

	processing.Normalize(in, normalized)
	showMatrix("Normalized", normalized)

	gx, gy := matrix.New(bounds), matrix.New(bounds)
	kernel.SobelDx.ConvoluteParallelized(normalized, gx)
	kernel.SobelDy.ConvoluteParallelized(normalized, gy)

	// Compute filtered directional
	filteredD, normFilteredD := matrix.New(bounds), matrix.New(bounds)
	kernel.FilteredDirectional(gx, gy, 4).ConvoluteParallelized(filteredD, filteredD)
	processing.Normalize(filteredD, normFilteredD)
	showMatrix("Filtered Directional", normFilteredD)

	// Compute segmented image
	segmented, normSegmented := matrix.New(bounds), matrix.New(bounds)
	kernel.Variance(filteredD).ConvoluteParallelized(normalized, segmented)
	processing.Normalize(segmented, normSegmented)
	showMatrix("Filtered Directional Std Dev.", normSegmented)

	// Compute binarized segmented image
	binarizedSegmented := matrix.New(bounds)
	processing.Binarize(normSegmented, binarizedSegmented)
	processing.BinarizeEnhancement(binarizedSegmented)
	showMatrix("Binarized Segmented", binarizedSegmented)

	// Binarize normalized image
	skeletonized := matrix.New(bounds)
	processing.Binarize(normalized, skeletonized)
	processing.BinarizeEnhancement(skeletonized)
	processing.Skeletonize(skeletonized)
	showMatrix("Skeletonized", skeletonized)

	// Run the whole thing again, the steps above are just intermediates
	// so we can have visibility on the algorithms
	result := extraction.DetectionResult(in)

	debug.DrawFeatures(img, result)
	showImage("Debug", img)
}

func main() {
	log.SetFlags(log.Flags() + log.Lshortfile)
	img, m := helpers.LoadImage("corpus/nist3.jpg")
	processImage(img, m)
}
