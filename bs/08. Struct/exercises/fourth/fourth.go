package main

import "fmt"

type Config struct {
	Timeout int
}

type App struct {
	Config *Config
}

func main() {
	c := Config{
		Timeout: 123,
	}
	// будет без инициализации (nil)
	// var a App

	// инициализировать правильно
	a := App{
		Config: &c,
	}

	fmt.Println(c, a)
}
