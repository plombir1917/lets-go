package structs

import (
	"errors"
	"fmt"
)

type Order struct {
	ID     int
	UserID int
	Amount int
	Status string
	Payed  bool
}

func (o *Order) Complete() error {
	if !o.Payed {
		return errors.New("Payed is false")
	}

	o.Status = "Closed"
	return nil
}

func Execute() {
	o := Order{Status: "Opened"}
	fmt.Println(o)

	err := o.Complete()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(o)

	o.Payed = true

	err = o.Complete()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(o)

}
