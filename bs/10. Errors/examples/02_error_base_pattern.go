package main

import (
	"errors"
	"log"
)

// базовый паттерн
func readFile(path string) error {
	if path == "" {
		return errors.New("empty path")
	}
	return nil
}

// использование
func basePatternInUse() {
	if err := readFile(""); err != nil {
		// ошибка это значение, а не исключение
		log.Fatal(err)

	}
}
