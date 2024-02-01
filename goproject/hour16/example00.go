package main

import (
	"fmt"
	"bytes"
)

func main () {
	var initn bytes.Buffer
	for i := 0; i < 20; i ++ {
		initn.WriteString("roqr")
	}
	fmt.Println(initn.String())
}