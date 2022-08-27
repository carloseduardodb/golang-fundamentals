package shapes

import (
	"math"
)

type Form interface {
	aria() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Square struct {
	radius float64
}

func (r Rectangle) Aria() float64 {
	return r.Width * r.Height
}

func (s Square) Aria() float64 {
	return math.Pi * math.Pow(s.radius, 2)
}
