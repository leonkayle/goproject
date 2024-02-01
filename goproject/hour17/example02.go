package main

import (
	"fmt"
	"flag"
)

func main () {
	s := flag.String("s", "Hello World", "String help text")
	flag.Parse()
	fmt.Println("value of s:", *s)
}