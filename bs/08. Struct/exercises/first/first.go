package main

import "fmt"

func main() {
	firstExercise()
}

type Counter struct {
	Value int
}

func (c Counter) Increment() {
	c.Value++
}
func (c *Counter) IncrementPtr() {
	c.Value++
}

func firstExercise() {

	c := Counter{
		Value: 1,
	}
	fmt.Println(c)

	c.Increment()
	fmt.Println(c)

	c.IncrementPtr()
	fmt.Println(c)

}
