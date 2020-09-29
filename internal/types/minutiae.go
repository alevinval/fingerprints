package types

type (
	MinutiaeType byte
	MinutiaeList []Minutiae
)

const (
	Termination MinutiaeType = iota
	Bifurcation
	Unknown
)

type Minutiae struct {
	X     int          `json:"x"`
	Y     int          `json:"y"`
	Angle float64      `json:"angle"`
	Type  MinutiaeType `json:"type"`
}
