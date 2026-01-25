package main

import (
	"fmt"
	"sync"
)

func printHello() {
	fmt.Println("Hello from printHello!")
}

func main() {

	go func() { fmt.Println("Hello from inline!") }()

	go printHello()

	fmt.Println("Hello from main!")

	waiting()
	wrongExample()
}

func waiting() {

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()
		fmt.Println("work")
	}()

	wg.Wait()
}

func wrongExample() {
	var result int

	go func() {
		result = 42
	}()

	fmt.Println(result)
}
