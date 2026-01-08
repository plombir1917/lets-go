package main

import "fmt"

func main() {
	slice := []int{1, 2, 3}
	fmt.Println(slice)
	// or
	sliceCreatedWithMake := make([]int, 3)
	fmt.Println(sliceCreatedWithMake)

	exampleWithAppend()

	rightSliceModify(slice)
	fmt.Println(slice)

	wrongSliceModify(slice)
	fmt.Println(slice)

	appendValue(slice)
	fmt.Println(slice)
}

func exampleWithAppend() {
	a := []int{1, 2, 3}
	// b := a
	b := make([]int, len(a))
	copy(b, a)

	b = append(b, 4)
	fmt.Println(a)
	fmt.Println(b)
}

func rightSliceModify(s []int) {
	s[0] = 100
}

//will create local variable without changing source slice
func wrongSliceModify(s []int) {
	s = append(s, 10)
}

func appendValue(s []int) []int {
	return append(s, 10)
}
