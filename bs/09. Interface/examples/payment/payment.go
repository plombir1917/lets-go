package payment

type PaymentMethod interface {
	Pay(usd int) int
	Cancel(id int)
}

type PaymentModule struct {
	payments      map[int]PaymentInfo
	paymentMethod PaymentMethod
}

// конструктор для paymentMethod
func NewPaymentModule(paymentMethod PaymentMethod) *PaymentModule {
	return &PaymentModule{
		payments:      make(map[int]PaymentInfo),
		paymentMethod: paymentMethod,
	}
}

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
	p.paymentMethod.Cancel(id)

	info, ok := p.payments[id]
	if !ok {
		return
	}
	info.Cancelled = true
	p.payments[id] = info
}

// принимает:
// ID операции
func (p *PaymentModule) Info(id int) PaymentInfo {
	info, ok := p.payments[id]
	if !ok {
		return PaymentInfo{}
	}
	return info
}

func (p *PaymentModule) AllInfo() map[int]PaymentInfo {
	// создаём локальную переменную, чтобы изменить исходные данные было невозможно
	tempMap := make(map[int]PaymentInfo, len(p.payments))

	for k, v := range p.payments {
		tempMap[k] = v
	}

	return tempMap
}
