package matching

import (
	"github.com/alevinval/fingerprints/internal/kernel"
	"github.com/alevinval/fingerprints/internal/matrix"
	"github.com/alevinval/fingerprints/internal/processing"
	"github.com/alevinval/fingerprints/internal/types"
)

func ExtractFeatures(in *matrix.M) types.MinutiaeList {
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

	return processing.ExtractFeatures(skeletonized, filteredD, binarizedSegmented)
}
