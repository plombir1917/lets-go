package main

import "fmt"

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

}
