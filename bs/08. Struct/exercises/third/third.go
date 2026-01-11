package main

import "fmt"

type Engine struct {
	Power int
}

type Car struct {
	Engine
	Brand string
}

func main() {

	e := Engine{
		Power: 150}

	c := Car{
		Engine: e,
	}

	fmt.Println(c.Power)
}
