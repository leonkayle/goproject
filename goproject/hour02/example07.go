package main

import (
	"fmt"
	"reflect"
)

func main () {
	var s string = "main"
	var i int = 13
	var f float64 = 5.68
	
	fmt.Println(reflect.TypeOf(s))
	fmt.Println(reflect.TypeOf(i))
	fmt.Println(reflect.TypeOf(f))
}