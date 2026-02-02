package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
)

var money = 1000
var bank = 0
var mtx = sync.Mutex{}

func payHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	for k, v := range r.Header {
		fmt.Println("k:", k, "v:", v)
	}
	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		msg := "Ошибка:" + err.Error()
		_, err := w.Write([]byte(msg))
		if err != nil {
			fmt.Println(err)
		}
		return
	}
	httpRequestBodyString := string(httpRequestBody)

	paymentAmount, err := strconv.Atoi(httpRequestBodyString)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		msg := "Ошибка:" + err.Error()
		_, err := w.Write([]byte(msg))
		if err != nil {
			fmt.Println(err)
		}
		return
	}

	mtx.Lock()
	if money >= (paymentAmount) {
		money -= paymentAmount
		msg := "Оплата прошла успешно:" + strconv.Itoa(money)
		_, err := w.Write([]byte(msg))
		if err != nil {
			fmt.Println(err)
		}
		mtx.Unlock()
		return
	}

}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		msg := "Ошибка:" + err.Error()
		_, err := w.Write([]byte(msg))
		if err != nil {
			fmt.Println(err)
		}
		return
	}

	httpRequestBodyString := string(httpRequestBody)

	saveAmount, err := strconv.Atoi(httpRequestBodyString)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		msg := "Ошибка:" + err.Error()
		_, err := w.Write([]byte(msg))
		if err != nil {
			fmt.Println(err)
		}
		return
	}

	mtx.Lock()
	if money > (saveAmount) {
		money -= saveAmount
		bank += saveAmount
		msg := "Деньги отложены. Состояние счёта:" + strconv.Itoa(bank)
		w.Write([]byte(msg))
		mtx.Unlock()
		return
	}

}

func main() {
	http.HandleFunc("/pay", payHandler)
	http.HandleFunc("/save", saveHandler)

	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		fmt.Println("Ошибка сервера:", err)
	}
}
