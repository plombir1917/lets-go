package main

import "fmt"

func calcTotalAmount(orderTotal float64, tip float64, paid float64) (float64, float64) {
	totalAmountDue := orderTotal + tip
	change := paid - totalAmountDue

	return totalAmountDue, change
}

func main() {
	totalCost, change := calcTotalAmount(8.00, 2.00, 10.00)
	fmt.Printf("Total cost is%.2f\nchange is %.2f", totalCost, change)
}
