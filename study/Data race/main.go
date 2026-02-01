package main

import (
	"fmt"
	"sync"
)

// var number int = 0
// var number atomic.Int64
var slice []int

var mtx sync.Mutex

func increase(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 1000; i++ {
		mtx.Lock()
		slice = append(slice, i)
		mtx.Unlock()
	}

}

func main() {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go increase(wg)
	wg.Add(1)
	go increase(wg)
	wg.Add(1)
	go increase(wg)

	wg.Add(1)
	go increase(wg)
	wg.Add(1)
	go increase(wg)
	wg.Add(1)
	go increase(wg)

	wg.Add(1)
	go increase(wg)
	wg.Add(1)
	go increase(wg)
	wg.Add(1)
	go increase(wg)

	wg.Wait()

	fmt.Println(len(slice))
}
