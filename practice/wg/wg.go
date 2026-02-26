package wg

import (
	"fmt"
	"sync"
)

func Execute() {
	s := []int{1, 2, 3}

	wg := &sync.WaitGroup{}

	for i := range s {

		wg.Go(func() {
			s[i] = 100
			fmt.Println("sss", s)

		})

	}

	wg.Wait()

	fmt.Println("fin", s)

}
