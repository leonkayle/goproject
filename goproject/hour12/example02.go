package main

import (
	"fmt"
	"time"
)

func slowFunc(c chan string) {
	time.Sleep(time.Second * 1)
	c <- "take message to chan."
}

func main () {
	c := make(chan string)
	go slowFunc(c)
	msg := <-c
	fmt.Println(msg)
}