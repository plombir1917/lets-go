package main

type PaymentMethod interface {
	Pay(usd int) int
	Cancel(id int)
}

type PaymentModule struct {
	paymentMethod PaymentMethod
}

// конструктор для paymentMethod
func NewPaymentModule(paymentMethod PaymentMethod) *PaymentModule {
	return &PaymentModule{
		paymentMethod: paymentMethod,
	}
}

// принимает:
// описание оплаты
// сумма операции
// ID проведённой операции
func (p PaymentModule) Pay(description string, usd int) int {
	id := p.paymentMethod.Pay(usd)

}

// принимает:
// ID операции
func (p PaymentModule) Cancel(id int) {}

// принимает:
// ID операции
func (p PaymentModule) Info() {}

func (p PaymentModule) AllInfo() {}
