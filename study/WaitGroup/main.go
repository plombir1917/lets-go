package main

import (
	"fmt"
	"sync"
	"time"
)

func postman(wg *sync.WaitGroup, text string) {
	defer wg.Done() // уменьшаем счётчик на 1

	for i := 1; i <= 3; i++ {
		fmt.Println("Я почтальон, вот текст газеты:", text, "в", i, "раз")
		time.Sleep(1 * time.Second)
	}

}

func main() {
	wg := &sync.WaitGroup{}

	wg.Add(1) // увеличиваем счётчик на 1
	go postman(wg, "Xyi")

	wg.Add(1) // увеличиваем счётчик на 1
	go postman(wg, "Pzd")

	wg.Add(1) // увеличиваем счётчик на 1
	go postman(wg, "Ssk")

	wg.Wait() // ожидаем выполнения горутин, пока счётчик станет 0

}
