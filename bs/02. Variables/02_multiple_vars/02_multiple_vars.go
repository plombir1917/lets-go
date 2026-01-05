package main

import "fmt"

// Declaration and initialization outside of function
var price = 2.50

func main() {
	// Declare and initialize with var explicit type
	var coffeeName string = "Espresso"

	// Type inferred, implicit
	var size = "Medium"

	// Shorter declaration and initialization. Possible only inside functions
	price := 2.50

	fmt.Println("Medium Espresso price is $2.50")
	fmt.Println(size, coffeeName, "price is", price)
	fmt.Printf("%s %s price is $%.2f\n", size, coffeeName, price)
	fmt.Println("Done")

}
