package main

import "fmt"

func calculatePriceAfterDiscount(price float64, discountRate float64) float64 {
	newPrice := price - (price * discountRate)
	return newPrice
}

func main() {
	var coffeePrice float64 = 5.00

	discount := 0.10
	fmt.Println("Basic price:", coffeePrice)

	coffeePrice = calculatePriceAfterDiscount(coffeePrice, discount)

	fmt.Println("Price with discount:", coffeePrice)
}
