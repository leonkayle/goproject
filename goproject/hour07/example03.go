package main

import "fmt"

type Music struct {
	Name string
	Bag int
}

func main () {
	mus := new(Music)
	mus.Name = "birthday song"
	mus.Bag = 12
	fmt.Println(mus)
}