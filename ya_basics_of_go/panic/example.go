package main

import "fmt"

func example() {
	myPanic()
}
func myPanic() {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("Возникла паника:", p)
		}
	}()
	panic("паникуем")
}
