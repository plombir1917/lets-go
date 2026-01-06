package main

import "fmt"

func getDrinkInfo(customerName string, drink string, price int) {
	info := "%s's favourite drink is %s and it's price is %v"
	fmt.Printf(info, customerName, drink, price)
}

func main() {
	getDrinkInfo("Zakhar", "Beer", 1488)
}
