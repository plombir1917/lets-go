package main

import (
	"errors"
	"fmt"
	"io/fs"
)

func errorUnpacking(err error) {
	var pathErr *fs.PathError

	if errors.As(err, &pathErr) {
		fmt.Println(pathErr.Path)
	}
}
