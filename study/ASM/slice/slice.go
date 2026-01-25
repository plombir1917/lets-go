package slice

import "fmt"

func Slice() {
	slice := []int{1, 2, 3}
	fmt.Println(slice)

	slice = append(slice, 4)
	fmt.Println(slice)
}

func SliceWithMake() {
	slice := make([]int, 1)
	fmt.Println(slice)

	slice = append(slice, 100, 100)
	fmt.Println(slice)
}
