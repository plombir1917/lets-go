package variables

import "fmt"

func Execute() {
	fmt.Println("___Variables___")

	var count int

	add(count)
	fmt.Println(count)

	addWithPtr(&count)
	fmt.Println(count)
}

func add(count int) {
	count++
}

func addWithPtr(count *int) {
	*count++
}
