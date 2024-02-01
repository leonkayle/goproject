package main

import "fmt"

func feedMe (portion int, eaten int) int{
	eaten = portion + eaten
	if eaten >= 5 {
		return eaten
	}
	return feedMe(portion, eaten)
}

func main () {
	fmt.Println(feedMe(1, 0))
}