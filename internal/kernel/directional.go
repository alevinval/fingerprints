package kernel

import (
	"math"

	"github.com/alevinval/fingerprints/internal/matrix"
)

type directional struct {
	Base
	gx, gy *matrix.M
}

func Directional(gx, gy *matrix.M) *directional {
	k := &directional{gx: gx, gy: gy}
	k.Base = Base{kernel: k}
	return k
}

func (k *directional) Offset() int {
	return 0
}

func (k *directional) Apply(_ *matrix.M, x, y int) float64 {
	dx := k.gx.At(x, y)
	dy := k.gy.At(x, y)
	ang := math.Atan2(dy, dx) + math.Pi/2
	return ang
}
