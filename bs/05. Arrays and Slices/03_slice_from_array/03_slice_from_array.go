package main

import "fmt"

func main() {
	desertMenu := [4]string{"Muffin", "Brownie", "Croissant", "Cookie"}

	fmt.Println((desertMenu))

	slice := desertMenu[1:2]

	fmt.Println(slice)

	slice = desertMenu[:] // all elements
	fmt.Println(slice)
}
