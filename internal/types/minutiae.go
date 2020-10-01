package types

import "fmt"

type (
	MinutiaeType byte
	MinutiaeList []Minutiae
)

const (
	Termination MinutiaeType = iota
	Bifurcation
	Pore
	Unknown
)

type Minutiae struct {
	X     int          `json:"x"`
	Y     int          `json:"y"`
	Angle float64      `json:"angle"`
	Type  MinutiaeType `json:"type"`
}

func (m Minutiae) String() string {
	return fmt.Sprintf("[%d,%d] Type=%v, Angle=%f", m.X, m.Y, m.Type, m.Angle)
}

func (t MinutiaeType) String() string {
	switch t {
	case Termination:
		return "t"
	case Bifurcation:
		return "b"
	default:
		return "n/a"
	}
}
