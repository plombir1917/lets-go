package main

import "fmt"

func main() {
	var coffeePrice = 4.50
	fmt.Println("Coffee price", coffeePrice)
	// Compile time (code you write): var coffeePrice = 4.50
	// Runtime (what machine sees): [some memory address] = 4.50

	fmt.Println("Memory address of price 4.50", &coffeePrice)

	coffeePrice = 5.00
	fmt.Println("Memory address of price 5.00", &coffeePrice)

	// pointerToCoffeePrice := &coffeePrice // same as next
	var pointerToCoffeePrice *float64 = &coffeePrice

	/*go to the memory location, where pointerToCoffeePrice points to
	and change value in this location*/
	*pointerToCoffeePrice = 5.50
	fmt.Println(coffeePrice)
}
