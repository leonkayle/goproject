package main

import (
	"fmt"
	"time"
)

func pinger(c chan string) {
	t := time.NewTicker(time.Second * 1)
	for {
		c <- "ping"
		<- t.C
	}
}

func main () {
	message := make(chan string)
	go pinger(message)
	for {
		msg := <- message
		fmt.Println(msg)
	}
}