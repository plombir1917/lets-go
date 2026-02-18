package slices

import "fmt"

func Execute() {
	fmt.Println("\n___Slices___")
	// s := []int{1, 2, 3}
	s := make([]int, 5)
	fmt.Println(s)

	changeSlice(s)
	fmt.Println(s)
}

func changeSlice(s []int) {
	for i := range s {
		s[i] = 1
	}
}
