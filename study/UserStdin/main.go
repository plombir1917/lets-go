package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Введите команду")

	if ok := scanner.Scan(); !ok {
		fmt.Println("Ошибка")
		return
	}
	text := scanner.Text()

	arrayFromText := strings.Fields(text)

	fmt.Println("text:", arrayFromText)
}
