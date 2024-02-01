package main

import "fmt"

type Weather struct {
	Name string
	Temperature float64
}

func main () {
	w := Weather {
		Name : "hot",
		Temperature : 17.66,
	}
	fmt.Printf("%+v", w)
}