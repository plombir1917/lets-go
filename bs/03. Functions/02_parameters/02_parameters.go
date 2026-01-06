package main

import "fmt"

func greet(name string) {
	name = "123"
	fmt.Printf("welcome to the coffee shop %v!", name)
}

func main() {
	greet("name")
}
