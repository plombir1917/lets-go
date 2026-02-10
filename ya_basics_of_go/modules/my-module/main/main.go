package main

import (
	"fmt"

	"github.com/plombir1917/moexample/math" // ваш пакет
)

// чёто она делает
func main() {
	if sum := math.Add(1, 2); sum != 3 {

		panic(fmt.Sprintf("sum expected to be 3; got %d", sum))
	}

	fmt.Println("Well done!")
}
