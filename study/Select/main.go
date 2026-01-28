package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	intCh := make(chan int)
	strCh := make(chan string)

	go func() {
		i := 1
		for {
			time.Sleep(100 * time.Millisecond)

			intCh <- i
			i++

		}
	}()

	go func() {
		i := 1
		for {
			time.Sleep(100 * time.Millisecond)

			strCh <- "hi" + strconv.Itoa(i)
			i++

		}
	}()

	// for {
	select {
	case number := <-intCh:
		fmt.Println(number)

	case str := <-strCh:
		fmt.Println(str)

	default:
		fmt.Println("Никакой канал не готов")

	}

	// }
}
