package extraction

import (
	"github.com/alevinval/fingerprints/internal/kernel"
	"github.com/alevinval/fingerprints/internal/matrix"
	"github.com/alevinval/fingerprints/internal/processing"
	"github.com/alevinval/fingerprints/internal/types"
)

func DetectionResult(in *matrix.M) *types.DetectionResult {
	bounds := in.Bounds()

	normalized := matrix.New(bounds)
	processing.Normalize(in, normalized)

	gx, gy := matrix.New(bounds), matrix.New(bounds)
	filteredD := matrix.New(bounds)
	kernel.SobelDx.ConvoluteParallelized(normalized, gx)
	kernel.SobelDy.ConvoluteParallelized(normalized, gy)
	kernel.FilteredDirectional(gx, gy, 4).ConvoluteParallelized(filteredD, filteredD)

	segmented := matrix.New(bounds)
	kernel.Variance(filteredD).ConvoluteParallelized(normalized, segmented)
	processing.Normalize(segmented, segmented)

	binarizedSegmented := matrix.New(bounds)
	processing.Binarize(segmented, binarizedSegmented)
	processing.BinarizeEnhancement(binarizedSegmented)

	skeletonized := matrix.New(bounds)
	processing.Binarize(normalized, skeletonized)
	processing.BinarizeEnhancement(skeletonized)
	processing.Skeletonize(skeletonized)

	minutia := Minutia(skeletonized, filteredD, binarizedSegmented)
	frame := Frame(binarizedSegmented)
	dr := &types.DetectionResult{
		Frame:   frame,
		Minutia: minutia,
	}
	return dr
}
