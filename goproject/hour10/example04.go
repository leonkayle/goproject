package main

import "fmt"

func Half(numberToHalf int) (int, error) {
	if numberToHalf%2 != 0 {
		return -1, fmt.Errorf("Cannot half %d", numberToHalf)
	}
	return numberToHalf, nil
}

func main () {
	t, err := Half(18)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Cannot half %d", t)
}