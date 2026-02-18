package conditions

import "fmt"

func Execute() {
	fmt.Println("\n___Conditions___")
	checkNumber(4)
	checkNumber(2)
	checkNumber(3)
}

func checkNumber(n int) {
	switch {
	case n%2 == 0 && n%4 == 0:
		fmt.Println("Одно сообщение")
	case n%2 == 0:
		fmt.Println("Другое сообщение")
	default:
		fmt.Println("Третье сообщение")
	}
}
