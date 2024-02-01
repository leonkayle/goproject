package main

import "fmt"

type Food struct {
	Name string
	Frozen bool
}

func main () {
	a := Food {
		Name : "fish",
		Frozen : false,
	}
	b := &a
	fmt.Printf("%+v\n", a)
	fmt.Printf("%p\n", &a)
	fmt.Printf("%+v\n", *b)
	fmt.Printf("%p\n", b)
}