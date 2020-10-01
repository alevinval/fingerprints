package types

type DetectionResult struct {
	Frame   Frame        `json:"frame"`
	Angle   float64      `json:"angle"`
	Minutia MinutiaeList `json:"minutia"`
}

func (dr *DetectionResult) RelativeMinutia() MinutiaeList {
	list := MinutiaeList{}

	for _, minutiae := range dr.Minutia {
		minutiae.X -= dr.Frame.Diagonal.Min.X
		minutiae.Y -= dr.Frame.Diagonal.Min.Y
		minutiae.Angle -= dr.Frame.Angle
		list = append(list, minutiae)
	}

	return list
}
