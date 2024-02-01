package main

import "fmt"

func echo (s string) {
	fmt.Println(s)
	return
}

func main () {
	s := "there are no signal."
	t := "off on the horizon."
	echo(s)
	echo(t)
}