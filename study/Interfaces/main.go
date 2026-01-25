package main

import (
	"interfaces/auto"
	"interfaces/payments"
	"interfaces/payments/methods"

	"github.com/k0kubun/pp"
)

func main() {
	auto.MainAuto()

	method := methods.NewCrypto()
	paymentModule := payments.NewPaymentModule(method)

	id := paymentModule.Pay("1я оплата", 1)
	paymentModule.Cancel(id)
	paymentModule.Pay("2я оплата", 2)
	paymentModule.Pay("3я оплата", 3)
	// paymentModule.Cancel(1)

	// paymentModule.Info(1)
	allInfo := paymentModule.AllInfo()

	pp.Println(allInfo)
	pp.Println(paymentModule.Info(id))
}
