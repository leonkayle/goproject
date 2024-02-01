package main

import "fmt"

func main () {
	var cheeses = make([]string, 3)
	cheeses[0] = "kata"
	cheeses[1] = "goodboy"
	cheeses[2] = "American"
	fmt.Println(cheeses)
	var smallCheeses = make([]string, 2)
	copy(smallCheeses, cheeses)
	fmt.Println(smallCheeses)
}