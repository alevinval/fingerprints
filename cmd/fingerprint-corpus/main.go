package main

import (
	_ "image/jpeg"
	"image/png"
	"log"
	"os"
	"path"

	"github.com/alevinval/fingerprints/internal/cmdhelper"
	"github.com/alevinval/fingerprints/internal/kernel"
	"github.com/alevinval/fingerprints/internal/matrix"
	"github.com/alevinval/fingerprints/internal/processing"
)

var outFolder = "out"

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
	kernel.SobelDx.ConvoluteParallelized(normalized, gx)
	kernel.SobelDy.ConvoluteParallelized(normalized, gy)

	// Compute filtered directional
	filteredD, normFilteredD := matrix.New(bounds), matrix.New(bounds)
	kernel.FilteredDirectional(gx, gy, 4).ConvoluteParallelized(filteredD, filteredD)
	processing.Normalize(filteredD, normFilteredD)
	showImage("Filtered Directional", normFilteredD)

	// Compute segmented image
	segmented, normSegmented := matrix.New(bounds), matrix.New(bounds)
	kernel.Variance(filteredD).Convolute(normalized, segmented)
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

	minutiaes := processing.ExtractFeatures(skeletonized, filteredD, binarizedSegmented)

	out := matrix.New(bounds)
	for _, minutiae := range minutiaes {
		log.Printf("found: %s", minutiae)
		out.Set(minutiae.X, minutiae.Y, 255.0)
	}

	showImage("Minutiaes", out)
}

func main() {
	log.SetFlags(log.Flags() + log.Lshortfile)
	img := cmdhelper.LoadImage("corpus/nist3.jpg")
	processImage(img)
}
