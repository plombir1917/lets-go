package main

import "fmt"

func main() {
	// firstExercise()
	// secondExercise()
	thirdExercise()
	// fourthExercise()
}

func firstExercise() {
	m := make(map[string]int)

	// добавление
	m["first"] = 1
	m["second"] = 2
	m["third"] = 3

	// обновление
	m["first"] = 10

	// удаление
	delete(m, "third")

	fmt.Println(m)
}

func secondExercise() {
	m := make(map[string]int)

	m["zero"] = 0

	// прямое чтение
	if m["zero"] == 0 {
		fmt.Println("Ловушка! Значение есть, и проверка неверна")
	}

	// value, ok
	v, ok := m["zero"]

	if ok {
		fmt.Println("Проверка верна! Явно убедились что значение есть, вот оно:", v)
	}
}

func thirdExercise() {
	// для проверки условия с nil
	// var nilM map[string]int
	emptlM := make(map[string]int)

	m := make(map[string]int)
	m["first"] = 1

	increment := func(m map[string]int, key string) {

		if m == nil {
			fmt.Println("Ошибка: передан пустой map")
			return
		}

		m[key]++

	}

	increment(emptlM, "first")
	fmt.Println(emptlM)

	increment(m, "first")
	fmt.Println(m)
}

func fourthExercise() {
	m := map[string]int{
		"first":  1,
		"second": 2,
		"third":  3,
		"fourth": 4,
		"fifth":  5,
	}

	fmt.Println(m)
	fmt.Println(m)
	fmt.Println(m)
	fmt.Println(m)
}
