package mapp

import "fmt"

func Map() {
	weather := map[int]int{
		1: 10,
		2: 13,
	}
	if weather[3] == 0 {
		fmt.Println(weather)

	}

	if _, ok := weather[3]; !ok {
		fmt.Println(ok)
	}

}
