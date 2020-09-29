package matching

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
	X     int
	Y     int
	Angle float64
	Type  MinutiaeType
}
