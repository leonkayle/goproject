package main

import "fmt"

type Movie struct {
	Name string
	Package float64
}

type Data struct {
	Sort int
	Detail Movie
}

func main () {
	d := Data {
		Sort : 1,
		Detail : Movie {
			Name : "Crisis",
			Package : 1.33,
		},
	}
	fmt.Printf("%+v\n", d)
}