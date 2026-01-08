package main

import "fmt"

func main() {
	var coffeeSizes [3]string // default element values

	fmt.Println(coffeeSizes)

	coffeeSizes[0] = "S"
	coffeeSizes[1] = "M"
	coffeeSizes[2] = "L"
	fmt.Println("First el is:", coffeeSizes[0])

	fmt.Println(coffeeSizes)

}
