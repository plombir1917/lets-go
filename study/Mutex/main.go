package main

import (
	"fmt"
	"sync"
	"time"
)

var likes int = 0
var mtx sync.RWMutex

func setLike(wg *sync.WaitGroup) {
	defer wg.Done()

	for range 100_000 {
		mtx.Lock()
		likes++
		mtx.Unlock()
	}
}

func getLike(wg *sync.WaitGroup) {
	defer wg.Done()

	for range 100_000 {
		mtx.RLock()
		_ = likes
		mtx.RUnlock()
	}
}

func main() {
	wg := &sync.WaitGroup{}
	initTime := time.Now()
	for range 10 {
		wg.Add(1)
		go setLike(wg)
	}

	for range 10 {
		wg.Add(1)
		go getLike(wg)
	}

	wg.Wait()

	fmt.Println(time.Since(initTime))
}
