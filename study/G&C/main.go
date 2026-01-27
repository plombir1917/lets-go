package main

import (
	"fmt"
	"time"
)

func main() {
	initTime := time.Now()

	coal := 0

	transferPoint := make(chan int)
	go mine(transferPoint, 1)
	go mine(transferPoint, 2)
	go mine(transferPoint, 3)

	coal += <-transferPoint
	coal += <-transferPoint
	coal += <-transferPoint

	fmt.Println("Добыто угля:", coal)
	fmt.Println("Прошло времени: ", time.Since(initTime))
}

func mine(transferPoint chan int, n int) {
	fmt.Println("Поход в шахту №", n, "начался...")
	time.Sleep(1 * time.Second)
	fmt.Println("Поход в шахту №", n, "закончился...")

	transferPoint <- 10
}
