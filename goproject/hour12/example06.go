package main

import (
	"fmt"
	"time"
)

func runFunc1(c chan string) {
	time.Sleep(time.Second * 3)
	c <- "in the way one"
}

func runFunc2(c chan string) {
	time.Sleep(time.Second * 2)
	c <- "in the way two"
}

func main () {
	c := make(chan string)
	d := make(chan string)
	go runFunc1(c)
	go runFunc2(d)
	select {
		case msg1 := <-c :
		fmt.Println(msg1)
		case msg2 := <-d :
		fmt.Println(msg2)
	}
}