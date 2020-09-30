package types

import "image"

type Frame struct {
	Horizontal image.Rectangle `json:"horizontal"`
	Vertical   image.Rectangle `json:"vertical"`
	Diagonal   image.Rectangle `json:"diagonal"`
	Angle      float64         `json:"angle"`
}
