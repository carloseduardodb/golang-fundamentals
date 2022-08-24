package main

import (
	"fmt"
	"math"
)

type form interface {
	aria() float64
}

func writeAria(f form) {
	fmt.Println(f.aria())
}

type rectangle struct {
	width  float64
	height float64
}

type square struct {
	radius float64
}

func (r rectangle) aria() float64 {
	return r.width * r.height
}

func (s square) aria() float64 {
	return math.Pi * math.Pow(s.radius, 2)
}

func main() {
	r := rectangle{10, 15}
	writeAria(r)

	c := square{10}
	writeAria(c)
}
