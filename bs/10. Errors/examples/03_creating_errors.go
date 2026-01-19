package main

import (
	"errors"
	"fmt"
)

// создание ошибок
func errorCreating() {

	// errors.New
	err1 := errors.New("something went wrong")

	// fmt.Errorf
	err2 := fmt.Errorf("something went wrong")

	fmt.Println(err1, err2)
}
