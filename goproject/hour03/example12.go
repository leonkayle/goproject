package main

import "fmt"

func showMemoryAddress (i *int) {
	fmt.Println(i)
	return
}

func main () {
	i := 7
	fmt.Println(&i)
	showMemoryAddress(&i)
}