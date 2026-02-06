package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func test(str string, a rune) (i int, is bool) {
	for i, v := range str {
		if v == a {
			return i, true
		}
	}
	return
}

func main() {
	fmt.Println(fact(5))
	fmt.Println(fib(5))

	iterator := generate(0)
	iterator()
	iterator()
	iterator()

	test, _ := area((circle))
	fmt.Println(test(4))
	panicingFunc()
}

func generate(x int) func() {
	return func() {
		fmt.Println(x)
		x += 2
	}
}

func fact(n int) int {
	if n == 0 {
		return 1
	} else {
		return n * fact(n-1)
	}

}

func fib(n int) int {
	switch {
	case n <= 1:
		return n
	default:
		return fib(n-1) + fib(n-2)
	}
}

type figures int

const (
	square   figures = iota // квадрат
	circle                  // круг
	triangle                // равносторонний треугольник
)

func area(f figures) (func(float64) float64, bool) {
	switch f {
	case square:
		return func(x float64) float64 { return x * x }, true
	case circle:
		return func(x float64) float64 { return 3.142 * x * x }, true
	case triangle:
		return func(x float64) float64 { return 0.433 * x * x }, true
	default:
		return nil, false
	}
}

func PrintAllFilesWithFilterClosure(path string, predicate func(string) bool) {
	// создаём переменную, содержащую функцию обхода
	// мы создаём её заранее, а не через оператор :=, чтобы замыкание могло сослаться на него
	var walk func(string)
	walk = func(path string) {
		// получаем список всех элементов в папке (и файлов, и директорий)
		files, err := os.ReadDir(path)
		if err != nil {
			fmt.Println("unable to get list of files", err)
			return
		}
		//  проходим по списку
		for _, f := range files {
			// получаем имя элемента
			// filepath.Join — функция, которая собирает путь к элементу с разделителями
			filename := filepath.Join(path, f.Name())
			// печатаем имя элемента, если путь к нему содержит filter, который получим из внешнего контекста
			if predicate(filename) {
				fmt.Println(filename)
			}
			// если элемент — директория, то вызываем для него рекурсивно ту же функцию
			if f.IsDir() {
				walk(filename)
			}
		}
	}
	// теперь вызовем функцию walk
	walk(path)
}

// containsDot возвращает все пути, содержащие точки
func containsDot(s string) bool {
	return strings.Contains(s, ".")
}

func panicingFunc() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("caught panic:", r)
		}
	}()

	panic("паника")
}
