package constants

import "fmt"

const (
	created = iota
	paid
	shipped
	delivered
)

func string() {
	fmt.Printf("Created: %v, Paid: %v, Shipped: %v, Delivered: %v\n", created, paid, shipped, delivered)
}

func Execute() {
	fmt.Println("___Constants___")
	string()
}
