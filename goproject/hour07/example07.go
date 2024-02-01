package main

import (
	"fmt"
	"reflect"
)

type Drink struct {
	Name string
	Ice bool
}

func main () {
	c := Drink {
		Name : "Cafortor",
		Ice : true,
	}
	d := Drink {
		Name : "wine",
		Ice : true,
	}
	if c == d {
		fmt.Println("c and d is the similiar struct.")
	} else {
		fmt.Println("the struct is different between c and d .")
	}
	fmt.Printf("the struct of c show the form is below: %+v\n", c)
	fmt.Println(reflect.TypeOf(c))
	fmt.Printf("the struct of d show the form is below: %+v\n", d)
	fmt.Println(reflect.TypeOf(d))
}