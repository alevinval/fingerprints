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

	gx, gy := matrix.New(bounds), matrix.New(bounds)
	kernel.SobelDx.ConvoluteParallelized(normalized, gx)
	kernel.SobelDy.ConvoluteParallelized(normalized, gy)

	filteredD := matrix.New(bounds)
	kernel.FilteredDirectional(gx, gy, 4).ConvoluteParallelized(filteredD, filteredD)

	segmented := matrix.New(bounds)
	kernel.Variance(filteredD).Convolute(normalized, segmented)
	processing.Normalize(segmented, segmented)

	binarizedSegmented := matrix.New(bounds)
	processing.Binarize(segmented, binarizedSegmented)
	processing.BinarizeEnhancement(binarizedSegmented)

	skeletonized := matrix.New(bounds)
	processing.Binarize(normalized, skeletonized)
	processing.BinarizeEnhancement(skeletonized)
	processing.Skeletonize(skeletonized)

	minutiaes := processing.ExtractFeatures(skeletonized, filteredD, binarizedSegmented)

	out := matrix.New(bounds)
	for _, minutiae := range minutiaes {
		log.Printf("Found minutiae at %d, %d", minutiae.X, minutiae.Y)
		log.Printf("Type=%v, Angle=%f", minutiae.Type, minutiae.Angle)
		out.Set(minutiae.X, minutiae.Y, 255.0)
	}

	showImage("Minutiaes", out)
}

func main() {
	original := cmdhelper.LoadImage("corpus/nist3.jpg")
	img := matrix.NewFromGray(original)
	processImage(img)
}
