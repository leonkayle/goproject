package main

import (
	"fmt"
	"errors"
)

func main () {
	err := errors.New("something is wrong.")
	if err != nil {
		fmt.Println(err)
	}
}