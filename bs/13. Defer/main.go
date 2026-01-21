package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("Я первая отложенная функция") // выполнится 4
	}()

	defer func() {
		fmt.Println("Я вторая отложенная функция") // выполнится 3
	}()

	fmt.Println("Я первая функция") // выполнится 1
	fmt.Println("Я вторая функция") // выполнится 2
}
