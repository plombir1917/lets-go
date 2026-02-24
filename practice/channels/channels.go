package channels

import (
	"fmt"
)

func Execute() {
	ch := make(chan int)

	go producer(ch)

	consumer(ch)

}

func producer(ch chan<- int) {
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
}

func consumer(ch <-chan int) {
	for value := range ch {
		fmt.Println(value)
	}
}
