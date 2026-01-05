package main

import "fmt"

func main() {
	name := "Zakhar"
	price := 2.99
	ready := true
	orderedCount := 4
	stockCount := 500

	fmt.Printf("%T %T %T %T %v", name, price, ready, orderedCount, stockCount)
}
