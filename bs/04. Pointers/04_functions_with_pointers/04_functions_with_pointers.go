package main

import "fmt"

func applyDiscount(price *float64, discountRate float64) {
	*price = *price - (*price * discountRate)
}

func main() {
	var coffeePrice float64 = 5.00
	discount := 0.10

	fmt.Println("Basic price:", coffeePrice)

	applyDiscount(&coffeePrice, discount)

	fmt.Println("Price with discount:", coffeePrice)
}
