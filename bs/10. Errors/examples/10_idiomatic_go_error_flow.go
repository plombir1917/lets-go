package main

import (
	"errors"
	"fmt"
)

type repo struct {
	ID int
}

func (r repo) Find(id int) (*User, error) {
	// код бизнес-логики
	return nil, errors.New("test")
}

func idiomaticErrorExample() error {
	_, err := repo.Find(repo{}, 1)
	if err != nil {
		return fmt.Errorf("get user: %w", err)
	}

	return nil
}
