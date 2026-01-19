package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("b == 0")
	}
	return a / b, nil
}

func calculate() error {
	_, err := divide(1, 0)
	if err != nil {
		return fmt.Errorf("error: %w", err)
	}
	return nil
}

type NotEnoughMoneyError struct {
	Balance int
	Amount  int
}

func (e NotEnoughMoneyError) Error() string {
	return fmt.Sprintf("not enough money! balance: %d, amount: %d", e.Balance, e.Amount)
}

func fourthExercise() {
	var target NotEnoughMoneyError

	var currentError error = NotEnoughMoneyError{
		Balance: 1,
		Amount:  1}

	if errors.As(currentError, &target) {
		fmt.Println("caught NotEnoughMoneyError:", target.Balance, target.Amount)
	}

}

func main() {
	if err := calculate(); err != nil {
		fmt.Println(err)
	}
	fourthExercise()
}
