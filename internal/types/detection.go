package types

type DetectionResult struct {
	X       int          `json:"x"`
	Y       int          `json:"y"`
	Angle   float64      `json:"angle"`
	Minutia MinutiaeList `json:"minutia"`
}
