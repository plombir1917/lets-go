package methods

import (
	"fmt"
	"math/rand"
)

type Bank struct{}

func NewBank() Bank {
	return Bank{}
}

func (b Bank) Pay(usd int) int {
	fmt.Println("БАНК")
	fmt.Println("ВОТ СТОЛЬКО ДЕНЕГ:", usd)

	return rand.Int()
}

func (b Bank) Cancel(id int) {
	fmt.Println("ОТМЕНА БАНКА:", id)
}
