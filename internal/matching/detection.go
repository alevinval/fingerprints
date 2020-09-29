package matching

import (
	"github.com/alevinval/fingerprints/internal/kernel"
	"github.com/alevinval/fingerprints/internal/matrix"
	"github.com/alevinval/fingerprints/internal/processing"
	"github.com/alevinval/fingerprints/internal/types"
)

func Detection(in *matrix.M) *types.DetectionResult {
	bounds := in.Bounds()

	normalized := matrix.New(bounds)
	processing.Normalize(in, normalized)

	gx, gy := matrix.New(bounds), matrix.New(bounds)
	kernel.SobelDx.ConvoluteParallelized(normalized, gx)
	kernel.SobelDy.ConvoluteParallelized(normalized, gy)

	filteredD := matrix.New(bounds)
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

	minutia := processing.ExtractMinutia(skeletonized, filteredD, binarizedSegmented)
	frame := processing.ExtractFrame(binarizedSegmented)
	dr := &types.DetectionResult{
		Frame:   frame,
		Minutia: minutia,
	}
	return dr
}
