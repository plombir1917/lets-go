package pointer

import "fmt"

type user struct {
	balance int
}

func (u *user) depositWithPtrReceiver(amount int) {
	u.balance += amount
}

func (u user) depositWithValueReceiver(amount int) {
	u.balance += amount
}

func Execute() {
	fmt.Println("\n___Pointer___")
	u := user{balance: 0}

	u.depositWithValueReceiver(1)
	fmt.Println(u)

	u.depositWithPtrReceiver(1)
	fmt.Println(u)
}
