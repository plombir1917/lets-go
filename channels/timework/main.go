package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomTimeWork() {
	time.Sleep(time.Duration(rand.Intn(100)) * time.Second)
}

func predictableTimeWork() {
	ch := make(chan int)

	go func() {
		randomTimeWork()
		close(ch)
	}()

	select {
	case <-ch:
		fmt.Println("Долгая ф-ия выполнилась")
	case <-time.After(3 * time.Second):
		fmt.Println("Таймер вышел")
	}
}

func main() {
	predictableTimeWork()
}
