package main

import "fmt"

type Alarm struct {
	Time string
	Sound string
}

func NewAlarm(time string) Alarm {
	t := Alarm {
		Time : time,
		Sound : "Go school",
	}
	return t
}

func main () {
	s := NewAlarm("07:00")
	fmt.Println(s)
}