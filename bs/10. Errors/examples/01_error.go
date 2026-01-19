package main

// error на самом деле
type error interface {
	Error() string
}
