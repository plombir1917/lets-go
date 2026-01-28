package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	transferPoint := make(chan int)

	go func() {
		iterations := 3 + rand.Intn(4)

		fmt.Println(iterations)

		for i := 1; i <= iterations; i++ {
			transferPoint <- 10
			time.Sleep(300 * time.Millisecond)
		}
		close(transferPoint)
	}()

	coal := 0
	// for {
	// 	v, ok := <-transferPoint

	// 	if !ok {
	// 		fmt.Println("Все данные прочитаны")
	// 		break
	// 	}

	// 	coal += v
	// 	fmt.Println(coal, ok)

	// }

	// идиома go
	// цикл автоматически завершится, когда канал будет закрыт
	for v := range transferPoint {
		coal += v
		fmt.Println(coal)
	}

	fmt.Println(coal)

}
