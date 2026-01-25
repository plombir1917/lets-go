package methods

type PaymentMethod interface {
	Pay(usd int) int
	Cancel(id int)
}
