package main

import "fmt"

func main() {
	coffee := "Espresso"
	pointer := &coffee

	fmt.Println(coffee)
	fmt.Println(&coffee)
	fmt.Printf("%p\n", pointer)

	coffeeCopy := coffee // value is copied to another location

	fmt.Println(coffeeCopy)
	fmt.Println(&coffeeCopy)
	fmt.Printf("%p", &coffeeCopy)
}
