package main

import "fmt"

func main() {
	// firstExercise()
	secondExercise()
	// thirdExercise()
	// fourthExercise()
}

func firstExercise() {
	array := [5]int{1, 2, 3, 4, 5}
	slice := array[1:4]
	fmt.Println(array)

	slice[0] = 100
	fmt.Println(array)

}

func secondExercise() {
	slice := make([]int, 0, 10)
	secondSlice := slice

	fmt.Println(secondSlice)

	secondSlice = append(secondSlice, 10)
	fmt.Println(secondSlice)

	fmt.Println(slice)
}

func thirdExercise() {
	cloneSlice := func(s []int) []int {
		// copiedSlice := make([]int, 0)
		// return append(copiedSlice, s...)

		// better
		copied := make([]int, len(s))
		copy(copied, s)
		copied[0] = 100

		return copied
	}

	s := []int{1, 2, 3}

	copiedS := cloneSlice(s)
	fmt.Println(s, copiedS)
}

func fourthExercise() {
	slice := make([]int, 0, 10)

	for i := 1; i <= 10; i++ {
		slice = append(slice, i)
		fmt.Println(len(slice), cap(slice))
	}
}
