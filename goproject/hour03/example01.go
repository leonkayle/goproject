package main

import "fmt"

func main () {
	var s string = "spring"
	var u string
	u = "summer"
	v, w := "autumn", 4
	fmt.Printf("the first season of a year is %s, the second's is %s, the third's is %s, the fourth's is %v", s, u, v, w)
}