package main

import (
	"time"
	"fmt"
)

func receiver(c chan string) {
	for msg := range c {
		fmt.Println(msg)
	}
}

func main () {
	c := make(chan string, 2)
	c <- "Hello"
	c <- "World"
	close(c)
	fmt.Println("Pushed two messages onto Channel with no receivers")
	time.Sleep(time.Second * 1)
	receiver(c)
}