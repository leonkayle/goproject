package main

import (
	"fmt"
	"math"
)

type Sphere struct {
	Radius float64
}

func (s *Sphere) SurfaceArea() float64 {
	return float64(4) * math.Pi * (s.Radius * s.Radius)
}

func (s *Sphere) Volume() float64 {
	rthree := s.Radius * s.Radius * s.Radius
	return (float64(4) / float64(3)) * math.Pi * rthree 
}

func main () {
	s := Sphere {
		Radius : 1.5,
	}
	fmt.Println("the SurfaceArea is",s.SurfaceArea())
	fmt.Println("the Volume is",s.Volume())
}