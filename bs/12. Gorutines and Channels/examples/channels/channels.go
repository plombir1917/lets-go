package main

import (
	"fmt"
	"time"
)

func printHello(ch chan int) {
	fmt.Println("Hello from printHello!")
	// посылает значение в канал
	ch <- 2
}

func main() {
	// инициализируем канал типа int, ёмкостью 2
	ch := make(chan int, 2)

	go func() {
		fmt.Println("Hello inline!")
		ch <- 1
	}()

	// вызываем функцию как горутину
	go printHello(ch)
	fmt.Println("Hello from main!")

	// Получаем первое значение из канала
	// и сохраним его в переменной, чтобы позже распечатать
	i := <-ch
	fmt.Println("Recieved", i)

	// Получаем второе значение из канала
	// и не сохраняем его, потому что не будем использовать
	<-ch

	safeCommunication()
	unbuffered()
	buffered()
	chanWithSelect()
}

func safeCommunication() {

	ch := make(chan int)

	fmt.Println(1)
	go func() {
		ch <- 42
	}()

	v := <-ch

	fmt.Println(v)

}

func unbuffered() {

	ch := make(chan int)

	// ch <- 1

	fmt.Println(ch)
}

func buffered() {

	ch := make(chan int, 3)

	close(ch)

	if _, ok := <-ch; ok {
		ch <- 1
		ch <- 2
		ch <- 3 // ok
		ch <- 4 // блокировка
	}

	for v := range ch {
		fmt.Println(v)
	}

	fmt.Println(<-ch)

}

func chanWithSelect() {

	ch := make(chan int)

	ch <- 1

	select {
	case v := <-ch:
		fmt.Println(v)
	case <-time.After(time.Second):
		fmt.Println("Timeout")
	}
}
