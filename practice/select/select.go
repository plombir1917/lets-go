package selecte

import (
	"fmt"
	"time"
)

func Execute() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go sendInFirst(ch1)
	go sendInSecond(ch2)

	select {
	case v := <-ch1:
		fmt.Println("first:", v)
	case v := <-ch2:
		fmt.Println("second:", v)
	}
}

func sendInFirst(ch chan<- int) {
	time.Sleep(3 * time.Second)
	ch <- 1
}

func sendInSecond(ch chan<- int) {
	time.Sleep(2 * time.Second)
	ch <- 2
}
