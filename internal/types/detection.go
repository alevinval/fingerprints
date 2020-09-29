package types

type DetectionResult struct {
	Frame   Frame        `json:"frame"`
	Angle   float64      `json:"angle"`
	Minutia MinutiaeList `json:"minutia"`
}
