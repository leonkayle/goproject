package main

import (
	"fmt"
	"time"
)

func slowFunc() {
	time.Sleep(time.Second * 2)
	fmt.Println("the road is aside.")
}

func main () {
	go slowFunc()
	fmt.Println("run away.")
	time.Sleep(time.Second * 3)
}