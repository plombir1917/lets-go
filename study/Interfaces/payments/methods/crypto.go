package methods

import (
	"fmt"
	"math/rand"
)

type Crypto struct {
}

func NewCrypto() Crypto {
	return Crypto{}
}

func (c Crypto) Pay(usd int) int {
	fmt.Println("Крипта")
	fmt.Println("Размер оплаты: ", usd)

	return rand.Int()
}

func (c Crypto) Cancel(id int) {
	fmt.Println("Крипта")
	fmt.Println("Отмена оплаты")
}
