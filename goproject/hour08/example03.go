package main

import "fmt"

type Triangle struct {
	Bottom float64
	Heigh float64
}

func (d *Triangle) area() float64 {
	return d.Bottom * d.Heigh / 2
}

func main () {
	d := Triangle {
		Bottom : 2.5,
		Heigh : 3,
	}
	fmt.Println("the TRiangle's area is", d.area())
}