package main

import (
	mapp "ASM/map"
	"ASM/slice"
	"fmt"
)

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	array(arr)

	fmt.Println(arr)

	arr[1]++

	for i := range arr {
		fmt.Println(i, "-", arr[i])
	}

	slice.Slice()
	slice.SliceWithMake()

	mapp.Map()
}

func array(arr [5]int) {
	arr[0] = 100
}
