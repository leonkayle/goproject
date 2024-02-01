package main

import (
	"fmt"
	"math"
	"strconv"
)

func main () {
	pi := strconv.FormatFloat(math.Pi, 'f', 4, 64)
	fmt.Println(pi)
}