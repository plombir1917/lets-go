package main

import (
	"fmt"
	"sync"
)

func work(id int, wg *sync.WaitGroup) {
	fmt.Println(id)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	go work(1, &wg)
	go work(2, &wg)

	wg.Wait()

	fmt.Println("DONE")
}
