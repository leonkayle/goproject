package main

import "fmt"

type Animal struct {
	Name string
	Color string
}

func main () {
	a := Animal {
		Name : "lion",
		Color : "brown",
	}
	fmt.Printf("%v\n", a)
	fmt.Printf("%+v\n", a)
}