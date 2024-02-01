package main

import "fmt"

func main () {
	q, p := "joek", "form"
	err := fmt.Errorf("erion %v %v hh.", q, p)
	if err != nil {
		fmt.Println(err)
	}
}