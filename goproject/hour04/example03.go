package main

import "fmt"

func sumNumbers(numbers...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main () {
	fmt.Println(sumNumbers(15, 73, 643))
}