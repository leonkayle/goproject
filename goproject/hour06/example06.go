package main

import "fmt"

func main () {
	var combine = make(map[string]int)
	combine["one"] = 1
	combine["two"] = 2
	combine["three"] = 3
	fmt.Println(combine["two"])
}