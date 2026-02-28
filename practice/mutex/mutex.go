package mutex

import (
	"sync"
)

func Execute() {
	int := 0
	mtx := &sync.Mutex{}

	for i := 0; i <= 1000; i++ {
		go func() {
			mtx.Lock()
			int++
			mtx.Unlock()
		}()
	}

}
