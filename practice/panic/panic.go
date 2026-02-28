package panic

import "fmt"

func Execute() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	callPanic()
}

func callPanic() {
	panic("ааааа бляяяя")
}
