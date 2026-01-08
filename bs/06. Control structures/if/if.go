package main

import "fmt"

// var x int = -1

func main() {
	// if x > 0 {
	// 	fmt.Println("positive")
	// }
	// if x < 0 {
	// 	fmt.Println("negative")
	// }
	// if x == 0 {
	// 	fmt.Println("zero")
	// }
	if x := 0; x > 0 {
		fmt.Println("positive")
	} else if x < 0 {
		fmt.Println("negative")
	} else if x == 0 {
		fmt.Println("zero")
	}

}
