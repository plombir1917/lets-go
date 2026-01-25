package main

import (
	"fmt"
)

type Reader interface {
	Read(p []byte) (n int, err error)
}

type File struct{}

type Network struct{}

func (n Network) Read(p []byte) (int, error) {
	return 0, nil
}

func (f File) Read(p []byte) (int, error) {
	return 0, nil
}

func main() {

	var r Reader
	fmt.Println(r == nil) // true

	var f *File = nil
	r = f
	fmt.Println(r == nil) // false

	r = File{} // никаких implements

	fmt.Println(r)

	v, ok := r.(File)
	// panic версия
	// v := r.(File)
	fmt.Println(v, ok)

	emptyInterfaceExample()

}

// плохо
type BigInterface interface {
	Read()
	Write()
	Close()
	Flush()
}

// хорошо
type Writer interface {
	Write()
}

// и т.д.

func process(r Writer) {
	r.Write()
}

type MyError interface{}

func foo() error {
	// var e *MyError = nil
	// e != nil
	// return e
	return nil
}

func typeSwitchExample() {
	var r Reader

	switch v := r.(type) {
	case File:
		fmt.Println("file", v)
	case Network:
		fmt.Println("Network")
	}
}

func emptyInterfaceExample() {
	// устаревающий паттерн
	// var v interface{}
	var v any
	fmt.Println(v)
}
