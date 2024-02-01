package main

import "fmt"

func main () {
	var cheeses = make(map[string]int)
	cheeses["one"] = 1
	cheeses["two"] = 2
	cheeses["three"] = 3
	fmt.Println(cheeses)
	delete(cheeses, "two")
	fmt.Println(cheeses)
}