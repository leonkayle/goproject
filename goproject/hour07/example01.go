package main

import "fmt"

type Movie struct {
	Name string
	Age int
}

func main () {
	cha := Movie {
		Name : "leon",
		Age : 18,
	}
	fmt.Println(cha)
}