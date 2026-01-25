package main

import (
	"fmt"
	"study/greeting"
	"study/user"

	"github.com/k0kubun/pp"
)

func main() {
	greeting.SayHello()
	greeting.SayBad()

	user := user.NewUser("Имя", 19)

	fmt.Println(user)

	pp.Println(user)
}
