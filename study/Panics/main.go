package main

import "fmt"

func main() {
	defer func() {
		r := recover()

		fmt.Println(r)
	}()

	slice := []int{1, 2, 3}
	b := slice[5]
	fmt.Println(b)
}
