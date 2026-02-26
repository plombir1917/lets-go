package race

import "fmt"

func Execute() {
	int := 0

	for i := 0; i <= 1000; i++ {
		go func() {
			int++
		}()
	}

	fmt.Println(int)
}
