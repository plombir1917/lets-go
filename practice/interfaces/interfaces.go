package interfaces

type PaymentProcessor interface {
	Process(amount int) error
}

type CardPayment struct{}

func (c CardPayment) Process(amount int) error {
	return nil
}

type CryptoPayment struct{}

func (c CryptoPayment) Process(amount int) error {
	return nil
}

func Execute() {
	card := CardPayment{}
	crypto := CryptoPayment{}

	pay(card)
	pay(crypto)
}

func pay(p PaymentProcessor) PaymentProcessor {
	return p
}
