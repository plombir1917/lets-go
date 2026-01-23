package methods

import (
	"fmt"
	"math/rand"
)

type Crypto struct{}

func (c *Crypto) Pay(usd int) int {
	fmt.Println("Оплата криптой")

	return rand.Int()
}

func (c *Crypto) Cancel(id int) {
	fmt.Println("Отмена оплаты криптой")
}
