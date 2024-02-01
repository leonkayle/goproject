package main

import "fmt"

func main () {
	var cheeses = make ( []string, 2 )
	cheeses[0] = "ada"
	cheeses[1] = "biology"
	fmt.Println(cheeses)
	ch := append(cheeses, "fire", "water")
	fmt.Println(ch)
}