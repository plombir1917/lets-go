package errors

import (
	"errors"
	"fmt"
)

type customError struct {
	message string
}

func (c customError) Error() string {
	return c.message
}

func Execute() {
	err := getCustomError()
	fmt.Println(errors.As(err, &customError{message: "123"}))
}

func getCustomError() error {
	return customError{message: "123"}
}
