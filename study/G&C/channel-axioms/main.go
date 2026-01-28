package main

import "fmt"

func main() {
	var ch chan int // nil channel

	// go func() {
	// 	ch <- 1 // вечная блокировка
	// }()

	// res := <-ch // вечная блокировка

	// fmt.Println(res)

	// создаём новый открытый канал
	channel := make(chan int)

	// закрываем канал
	close(channel)

	// читаем значение из закрытого канала
	v, ok := <-channel
	fmt.Println(v, ok) // 0 - default значение для типа канала; ok - дефолтно ли полученное значение

	// запись значения в закрытый канал
	channel <- 1 // - паника

	// закрываем закрытый канал
	close(channel) // - паника

	fmt.Println(ch)
}
