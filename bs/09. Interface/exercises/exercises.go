package main

import "fmt"

type Speaker interface {
	Speak() string
}

type Human struct{}

func (h Human) Speak() string {
	return "Hello"
}

type Robot struct{}

func (r Robot) Speak() string {
	return "Hello"
}

func main() {

	firstExercise()
	secondExercise()
	e := thirdExercise()
	fmt.Println(e == nil)

	var h Human
	var r Robot
	fourthExercise(h)
	fourthExercise(r)
}

func firstExercise() {
	var h Human
	var r Robot

	h.Speak()
	r.Speak()
}

func secondExercise() {

	var h Human

	say := func(s Speaker) {
		fmt.Println(s.Speak())
	}

	say(h)

}

type MyError struct{}

func (e *MyError) Error() string {
	return "my error"
}

func thirdExercise() error {

	var e *MyError = nil
	return e
}

func fourthExercise(s Speaker) {
	switch v := s.(type) {
	case Human:
		fmt.Println("Human")
	case Robot:
		fmt.Println("Robot")
	default:
		fmt.Println(v)
	}
}
