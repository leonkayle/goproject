package main

import "fmt"

type Movie struct {
	Name string
	OpenSort int
}

func main () {
	var movies Movie
	fmt.Printf("%+v\n",movies)
	movies.Name = "coco"
	movies.OpenSort = 13
	fmt.Printf("%+v\n",movies)
}