package main

import "fmt"

func main() {
	firstMap := map[string]int{
		"first": 1,
	}
	fmt.Println(firstMap)

	// or

	secondMap := make(map[string]int)
	fmt.Println(secondMap)

	// mostPopularError()
	// operations()
	// read()
	// checkValue()
	// mapAndFunсtions()
	// structAsMapKey()
	mapInStruct()
}

func mostPopularError() {
	// неправильно
	// var m map[string]int
	// m["a"] = 10

	// правильно
	rightM := make(map[string]int)
	fmt.Println(rightM)
	rightM["a"] = 10
	fmt.Println(rightM)
}

func operations() {
	m := make(map[string]int)
	fmt.Println(m)

	m["first"] = 1
	fmt.Println(m)

	m["first"] = 10
	fmt.Println(m)

	delete(m, "first")
	fmt.Println(m)
}

func read() {
	m := make(map[string]int)
	m["first"] = 1

	v := m["first"]

	fmt.Println(v)
}

func checkValue() {
	m := make(map[string]int)
	m["first"] = 1

	v, ok := m["first"]
	if ok {
		fmt.Println(v)
	}
}

func mapAndFunсtions() {
	addValue := func(m map[string]int) {
		m["new"] = 1
	}

	m := make(map[string]int)
	fmt.Println(m)

	addValue(m)
	// изменения видны снаружи
	// нет append как в slice
	fmt.Println(m)
}

func structAsMapKey() {
	type Key struct {
		ID int
	}

	m := map[Key]string{}
	fmt.Println(m)
}

func mapInStruct() {
	type User struct {
		Scores map[string]int
	}

	// неправильно
	// var u User
	// u.Scores["first"] = 1 // panic

	// правильно
	u := User{
		Scores: make(map[string]int),
	}
	u.Scores["first"] = 1
	fmt.Println(u)
}
