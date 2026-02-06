package main

import (
	"fmt"
)

type Person struct {
	// теги
	Name  string `json:"Имя" example:"Alex"`
	Email string `json:"Почта" example:"мыло"`
	Age   int    `json:"-"`
}

func main() {
	p := Person{
		Name:  "Лёха",
		Email: "мыло", // ЗАПЯТАЯ СТАВИТСЯ НА ПОСЛЕДНЕМ СВОЙСТВЕ ВСЕГДА!!
	}

	secondPerson := NewPerson("Имя", "мыло", 12)

	fmt.Println(p, secondPerson, p.Name)
}

// Аналог конструктора
func NewPerson(name, email string, Age int) Person {
	return Person{name, email, Age}
}
