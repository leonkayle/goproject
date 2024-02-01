package main

import (
	"fmt"
	"strconv"
)

type Television struct {
	Time string
	Name string
	Show bool
	Much int
}

func (t *Television) show() string {
	r := strconv.Itoa(t.Much)
	w := strconv.FormatBool(t.Show)
	return t.Time + " " + t.Name + " " + r + " " + w
}

func main () {
	x := Television {
		Time : "07:00",
		Name : "Escape from cell",
		Show : true,
		Much : 2,
	}
	fmt.Println(x.show())
}