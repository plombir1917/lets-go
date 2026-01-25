package greeting

import "fmt"

func SayHello() {
	fmt.Println("Hello, golang packages!")

	i := GiveMeInt()
	fmt.Println("Int from SayHello:", i)
}
