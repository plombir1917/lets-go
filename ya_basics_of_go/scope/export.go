package main

import "fmt"

const Email = "email@example.ru"

var support string

func SetSupport(s string) {
	support = s
}

func GetContact() string {
	return fmt.Sprintf("%s <%s>", support, Email)
}