package main

import "fmt"

// обёртывание ошибок
func errorWrapping(err error) {
	err = fmt.Errorf("failed: %w", err)
	fmt.Println(err)
}
