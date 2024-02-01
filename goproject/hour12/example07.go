package main

import (
	"fmt"
	"time"
)

func sender(c chan string) {
	t := time.NewTicker(time.Millisecond * 500)
	for {
		c <- "I'm sending a message"
		<- t.C
	}
}

func main () {
	message := make(chan string)
	stop := make(chan bool)
	go sender(message)
	go func () {
		time.Sleep(time.Second * 1)
		fmt.Println("Time's up!")
		stop <- true
	} ()
	for {
		select {
			case <- stop :
			return
			case messages := <- message :
			fmt.Println(messages)
		}
	}
}