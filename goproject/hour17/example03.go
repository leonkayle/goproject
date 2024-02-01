package main

import (
	"fmt"
	"flag"
)

func main () {
	s := flag.String("s", "hello world", "String help text")
	i := flag.Int("i", 1, "Int help text")
	b := flag.Bool("b", true, "Bool help text")
	flag.Parse()
	fmt.Printf("%v\n", *s)
	fmt.Printf("%v\n", *i)
	fmt.Printf("%v\n", *b)
}