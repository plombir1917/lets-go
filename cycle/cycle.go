package main

import "fmt"

func main() {
	array := [3]int{1, 2, 3}
	for arrayIndex, arrayValue := range array {
		fmt.Printf("array[%d]: %d\n", arrayIndex, arrayValue)
	}
}
