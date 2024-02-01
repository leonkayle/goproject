package main

import "fmt"

func main () {
	defer fmt.Println("show me the programme is one.")
	defer fmt.Println("show me the programme is two.")
	fmt.Println("give your gaze at me.")
}