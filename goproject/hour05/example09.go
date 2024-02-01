package main

import "fmt"

func main () {
	arr := []int{1, 2, 3, 4}
	for k, v := range arr {
		fmt.Println("the key of sort is ", k)
		fmt.Println("the value of sort is ", v)
	}
}