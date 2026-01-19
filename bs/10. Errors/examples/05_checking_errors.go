package main

import (
	"errors"
	"io/fs"
)

func errorChecking(err error) {
	// плохо
	if err.Error() == "not found" {
	}

	// хорошо
	if errors.Is(err, fs.ErrNotExist) {
	}
}
