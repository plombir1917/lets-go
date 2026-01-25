package payments

import (
	"fmt"
	"interfaces/payments/methods"
	"maps"
)

type PaymentModule struct {
	payments      map[int]PaymentInfo
	paymentMethod methods.PaymentMethod
}

// constructor
func NewPaymentModule(paymentModule methods.PaymentMethod) *PaymentModule {
	return &PaymentModule{
		payments:      make(map[int]PaymentInfo),
		paymentMethod: paymentModule,
	}
}

// METHODS:
func (p *PaymentModule) Pay(description string, usd int) int {
	id := p.paymentMethod.Pay(usd)
	info := PaymentInfo{
		Description: description,
		Usd:         usd,
		Cancelled:   false,
	}

	p.payments[id] = info

	return id
}

func (p *PaymentModule) Cancel(id int) {
	info, ok := p.payments[id]

	if !ok {
		fmt.Println("No payment with id: ", id)
		return
	}

	p.paymentMethod.Cancel(id)

	info.Cancelled = true

	p.payments[id] = info
}

func (p *PaymentModule) Info(id int) PaymentInfo {
	info, ok := p.payments[id]

	if !ok {
		return PaymentInfo{}
	}

	return info
}

func (p *PaymentModule) AllInfo() map[int]PaymentInfo {
	tempMap := make(map[int]PaymentInfo, len(p.payments))
	maps.Copy(tempMap, p.payments)

	return tempMap
}
