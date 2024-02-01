package main

import (
	"fmt"
	"errors"
)

type Robot interface {
	PowerOn() error
}

type T850 struct {
	Name string
}

func (t *T850) PowerOn() error {
	return nil
}

type R2D2 struct {
	Broken bool
}

func (r *R2D2) PowerOn() error {
	if r.Broken {
		return errors.New("the R2D2 is Broken.")
	}
	return nil
}

func Boot(r Robot) error {
	return r.PowerOn()
}

func main () {
	a := T850 {
		Name : "T850",
	}
	b := R2D2 {
		Broken :true,
	}
	fmt.Println(Boot(&a))
	fmt.Println(Boot(&b))
}