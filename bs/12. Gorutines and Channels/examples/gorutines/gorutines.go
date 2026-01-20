package main

import "fmt"

func printHello() {
	fmt.Println("Hello from printHello!")
}

func main() {

	go func() { fmt.Println("Hello from inline!") }()
	go printHello()

	fmt.Println("Hello from main!")
}
