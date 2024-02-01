package main

import (
	"fmt"
	"bytes"
)

func main () {
	var buffer bytes.Buffer
	for i:=0;i<50;i++{
		buffer.WriteString("fi")
	}
	fmt.Println(buffer.String())
}