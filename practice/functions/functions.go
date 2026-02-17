package functions

import "fmt"

func Execute() {
	fmt.Println("___Functions___")
	s := []int{1, 2, 3}
	fmt.Println(multipleReturn(s))
}

func multipleReturn(slice []int) (sum int, avg float32, err error) {
	if len(slice) == 0 {
		return
	}

	for _, v := range slice {
		sum += v
	}

	avg = float32(sum / len(slice))

	return
}
