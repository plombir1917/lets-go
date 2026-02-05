package main

import "fmt"

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
