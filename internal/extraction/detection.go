package extraction

import (
	"github.com/alevinval/fingerprints/internal/kernel"
	"github.com/alevinval/fingerprints/internal/matrix"
	"github.com/alevinval/fingerprints/internal/processing"
	"github.com/alevinval/fingerprints/internal/types"
)

// DetectionResult applies all the available image processing
// algorithms to finally extract a list of minutiae. Reference
// frame is detected too.
func DetectionResult(in *matrix.M) *types.DetectionResult {
	bounds := in.Bounds()

	inMeta := processing.Metadata(in)
	normalized := matrix.New(bounds)
	processing.Normalize(in, normalized, inMeta)
	normalizedMeta := processing.Metadata(normalized)

	gx, gy := matrix.New(bounds), matrix.New(bounds)
	filteredD := matrix.New(bounds)
	kernel.SobelDx.ConvoluteParallelized(normalized, gx)
	kernel.SobelDy.ConvoluteParallelized(normalized, gy)
	kernel.FilteredDirectional(gx, gy, 4).ConvoluteParallelized(filteredD, filteredD)

	segmented := matrix.New(bounds)
	segmentedMeta := processing.Metadata(in)
	kernel.Variance(filteredD).ConvoluteParallelized(normalized, segmented)
	processing.Normalize(segmented, segmented, segmentedMeta)

	binarizedSegmented := matrix.New(bounds)
	processing.Binarize(segmented, binarizedSegmented, segmentedMeta)
	processing.BinarizeEnhancement(binarizedSegmented)

	skeletonized := matrix.New(bounds)
	processing.Binarize(normalized, skeletonized, normalizedMeta)
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
