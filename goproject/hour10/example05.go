package main

import (
	"fmt"
)

func main () {
	fmt.Println("this is executed.")
	panic("Oh no.I can do no more.Goodbye.")
	fmt.Println("this is not executed.")
}