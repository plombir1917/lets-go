package gorutine

import (
	"fmt"
	"sync"
)

func Execute() {
	wg := &sync.WaitGroup{}
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func() {
			fmt.Println(i)
			wg.Done()
		}()

	}
	wg.Wait()
}
