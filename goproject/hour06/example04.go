package main

import "fmt"

func main () {
	var cheeses = make ([]string ,3)
	cheeses[0] = "this"
	cheeses[1] = "is"
	cheeses[2] = "example"
	ch := append(cheeses[:1], cheeses[2:]...)
	fmt.Println(ch)
}