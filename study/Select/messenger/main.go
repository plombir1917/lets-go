package main

import (
	"fmt"
	"time"
)

type Message struct {
	Author string
	Text   string
}

func main() {
	messageChan1 := make(chan Message)
	messageChan2 := make(chan Message)

	go func() {
		for {
			messageChan1 <- Message{
				Author: "Кирюха",
				Text:   "Привет",
			}

			time.Sleep(100 * time.Millisecond)
		}
	}()

	go func() {
		for {
			messageChan2 <- Message{
				Author: "Арс",
				Text:   "Селям",
			}

			time.Sleep(100 * time.Millisecond)
		}
	}()

	for {
		select {
		case message1 := <-messageChan1:
			fmt.Println(message1)

		case message2 := <-messageChan2:
			fmt.Println(message2)

		default:
			fmt.Println("Сообщений нет")
		}

	}
}
