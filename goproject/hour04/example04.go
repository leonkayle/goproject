package main

import "fmt"

func sayHi () (x, y string) {
	x = "Hello"
	y = "World"
	return
}

func main () {
	fmt.Println(sayHi())
}