package main

import "fmt"

var x int = 0

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// while style
	for x < 10 {
		x++
	}

	// infinity loop
	// for {
	// 	fmt.Println("loop")
	// }

	printEveryNumberFromOneToHundred()
	printEvenNumbers()
	sumFromOneToHundred()
}

func printEveryNumberFromOneToHundred() {
	for i := 1; i <= 100; i++ {
		fmt.Println(i)
	}
}

func printEvenNumbers() {
	for i := 2; i <= 100; i += 2 {
		fmt.Println(i)
	}

	// better
	/*
	   func printEveryHonestNumberFromOneToHundred() {
	   	for i := 1; i <= 100; i++ {
	   		if i%2 == 0 {
	   			fmt.Println(i)
	   		}
	   	}
	*/
}

func sumFromOneToHundred() int {
	sum := 0

	for i := 1; i <= 100; i++ {
		sum += i
	}

	return sum
}
