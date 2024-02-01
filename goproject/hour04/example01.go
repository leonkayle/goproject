package main

import "fmt"

func isEven (i int) bool {
	if i%2 == 0 {
		return true
	}
	return false
}

func main () {
	i := 79
	b := isEven(i)
	if b == true {
		fmt.Println("the input number is odd.")
	} else {
		fmt.Println("the input number is not odd.")
	}
}