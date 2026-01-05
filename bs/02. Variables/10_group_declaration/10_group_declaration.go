package main

import "fmt"

func main() {
	var first, second = 1, 2

	fmt.Println(first, second)
	var (
		customerName string = "name"
		tableNum     int    = 2
		isReadyToPay bool   = true
	)

	fmt.Printf("customerName is %s\ntableNum is %d\nis he ready to pay: %v", customerName, tableNum, isReadyToPay)
}
