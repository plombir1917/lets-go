package main

import (
	"errors"
	"fmt"
	"strings"
)

func Mul(v any, b int) any {
	switch vWithType := v.(type) {
	case int:
		return (vWithType * b)
	case string:
		return strings.Repeat(vWithType, b)
	case fmt.Stringer:
		return strings.Repeat(vWithType.String(), b)
	default:
		return errors.New("unknown type")
	}
}

func main() {
	fmt.Println(Mul("string", 4))
	(Mul(fmt.Sprint(), 4))

}
