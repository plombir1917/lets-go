package main

import "fmt"

func writer() <-chan int {
	ch := make(chan int)

	go func() {
		for range 2 {
			ch <- 1
		}

		close(ch)
	}()

	return ch
}

func main() {

	ch := writer()

	for v := range ch {
		fmt.Println(v)
	}
}
