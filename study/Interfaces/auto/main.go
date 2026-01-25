package auto

import "fmt"

type Auto interface {
	StepOnGas()
	StepOnBrake()
}

type BMW struct{}

func (b BMW) StepOnGas() {
	fmt.Println("BMW GAS")
}

func (b BMW) StepOnBrake() {
	fmt.Println("BMW BRAKE")
}

type Zhiga struct{}

func (z Zhiga) StepOnGas() {
	fmt.Println("Zhiga gas")
}

func (z Zhiga) StepOnBrake() {
	fmt.Println("Zhiga brake")
}

func ride(auto Auto) {
	fmt.Println("Я ВОДИЛА")
	auto.StepOnGas()
}

func stop(auto Auto) {
	auto.StepOnBrake()
}

func MainAuto() {
	bmw := BMW{}
	bmw.StepOnGas()

	zhiga := Zhiga{}
	zhiga.StepOnGas()

	ride(bmw)
	ride(zhiga)

	stop(bmw)
	stop(zhiga)
}
