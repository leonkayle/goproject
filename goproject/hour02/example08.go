package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main () {
	var b bool = true
	fmt.Println(reflect.TypeOf(b))
	t := strconv.FormatBool(b)
	fmt.Println(reflect.TypeOf(t))
	fmt.Println(t)
}