package main

import "fmt"


func main() {
	SetSupport("Служба поддержки")
	fmt.Println(GetContact())
	fmt.Println("Email: ", Email)
}