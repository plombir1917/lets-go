package rwmutex

import (
	"fmt"
	"sync"
)

type WithMap struct {
	field map[string]int
}

func Execute() {
	var example = WithMap{field: make(map[string]int)}

	wg := &sync.WaitGroup{}

	rwmtx := &sync.RWMutex{}

	for range 1000 {
		wg.Go(func() {
			rwmtx.RLock()
			fmt.Print(example.field)
			rwmtx.RUnlock()
		})
	}

	for i := range 100 {
		wg.Go(func() {
			rwmtx.Lock()
			example.field["key"+fmt.Sprint(i)] = i
			rwmtx.Unlock()
		})
	}

	wg.Wait()

	fmt.Print("fin", example)
}
