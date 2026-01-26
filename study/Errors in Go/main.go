package main

import (
	"errors"
	"fmt"

	"github.com/k0kubun/pp"
)

type User struct {
	Name     string
	Ballance int
}

func Pay(user *User, usd int) (*User, error) {
	if user.Ballance-usd < 0 {
		return user, errors.New("balance below zero")
	}
	user.Ballance -= usd
	return user, nil
}

func main() {
	user := User{
		Name:     "Имя",
		Ballance: 10,
	}

	pp.Println(user)

	_, err := Pay(&user, 20)

	if err != nil {
		fmt.Println(err)
	}
	pp.Println(user)

}
