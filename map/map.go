package main

import "fmt"

func main() {
	type MyMap map[string]string

	var m1 MyMap

	m1 = make(MyMap, 5)

	m1["foo"] = "bar"

	fmt.Println(m1)
}
