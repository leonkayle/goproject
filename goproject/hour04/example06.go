package main

import "fmt"

func anotherFunction(f func() string) string {
	return f()
}

func main () {
	fn := func() string {
		return "function canceled"
	}
	fmt.Println(anotherFunction(fn))
}