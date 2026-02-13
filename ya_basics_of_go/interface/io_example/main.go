package main

import (
	"interface-example/hashbyte"
	"interface-example/randbyte"
)

type MyInterface interface {
	GetArrayOfInsertedArrayofInts() [][]int
}

func main() {
	// standart.Execute()
	randbyte.Execute()
	hashbyte.Execute()
}
