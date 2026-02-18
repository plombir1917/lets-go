package arrays

import "fmt"

func Execute() {
	fmt.Println("\n___Arrays___")
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)

	changeArray(arr)
	fmt.Println(arr)
}

func changeArray(arr [5]int) {
	for i := range arr {
		arr[i] = 1
	}
}
