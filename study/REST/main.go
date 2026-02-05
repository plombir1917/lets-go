package main

import (
	"fmt"
	"rest/http"
	"rest/todo"
)

func main() {
	todoList := todo.NewList()
	httpHandlers := http.NewHTTPHandler(todoList)
	httpServer := http.NewHTTPServer(httpHandlers)

	if err := httpServer.StartServer(); err != nil {
		fmt.Println(err)
	}
}
