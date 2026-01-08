package main

import "fmt"

var x int = -1

func main() {
	switch {
	case x > 0:
		fmt.Println("positive")
	case x < 0:
		fmt.Println("negative")
	default:
		fmt.Println("zero")
	}

}

func describe(v interface{}) {
	switch t := v.(type) {
	case int:
		fmt.Println("int", t)
	case string:
		fmt.Println("string", t)
	default:
		fmt.Println("unknown", t)
	}
}
