package main

import "fmt"

func square(n int, ch chan int) {
	ch <- n * n
}

func main() {
	ch := make(chan int) // initialize channel of type int

	go square(5, ch) // starts gorutine

	result := <-ch // stops main, waiting for channel to get a value

	fmt.Println(result)

}
