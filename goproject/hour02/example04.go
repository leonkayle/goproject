package main

import "fmt"

func addition(x int, y int) int {
	return x + y
}

func main () {
	s := "twenty"
	fmt.Println(addition(5, s))
}