package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	str := "Hello world"
	b := []byte(str)

	_, err := w.Write(b)
	if err != nil {
		fmt.Println("Ошибка:", err.Error())
	} else {
		fmt.Println("Корректно")
	}
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Сервер запущен")
	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		fmt.Println("Ошибка")
	}

	fmt.Println("Программа завершена")

}
