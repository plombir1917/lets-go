package main

import "fmt"

func main() {
	born := 2000
	if (born >= 1846) && (born <= 1964) {
		fmt.Println("Привет, бумер")
	} else if born >= 1965 && born <= 1980 {
		fmt.Println("Привет, Х")
	} else if born >= 1981 && born <= 1996 {
		fmt.Println("Привет, миллениал")
	} else if born >= 1997 && born <= 2012 {
		fmt.Println("Привет, зумер")
	} else {
		fmt.Println("Привет, альфа")
	}
}
