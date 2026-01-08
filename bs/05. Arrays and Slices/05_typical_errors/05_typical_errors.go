package main

import "fmt"

func main() {
	firstError()
	secondError()
	forWithSliceAndRange()
}

func firstError() {
	// var s []int
	// s[0] = 10 panic, because 's' is empty slice
	s := make([]int, 1)
	fmt.Println(s)
}

func secondError() {
	s := make([]int, 0, 100)
	// for i := 0; i < len(s); i++ { slice has no len yet
	for i := 0; i < 100; i++ {
		s = append(s, i)
	}
	fmt.Println(s)
}

func forWithSliceAndRange() {
	s := make([]int, 10)

	for i, v := range s {
		fmt.Println(i, v)
	}

	// if index not needed
	for _, v := range s {
		fmt.Println(v)
	}
}
