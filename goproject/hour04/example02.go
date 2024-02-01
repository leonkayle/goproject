package main

import "fmt"

func getPrize () (int, string) {
	i := 56
	s := "the value"
	return i, s
}

func main () {
	quantity, prize := getPrize()
	fmt.Printf("output data is blow as %v, %s", quantity, prize)
}