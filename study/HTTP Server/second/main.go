package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

type Payment struct {
	Description string `json:"description"`
	USD         int    `json:"usd"`
	FullName    string `json:"fullName"`
	Address     string `json:"address"`
}

func (p Payment) Println() {
	fmt.Println("description:", p.Description)
	fmt.Println("usd:", p.USD)
	fmt.Println("fullName:", p.FullName)
	fmt.Println("address:", p.Address)
}

var mtx = sync.Mutex{}
var paymentHistory = make([]Payment, 0)
var money = 1000

func payHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var payment Payment
	if err := json.Unmarshal(httpRequestBody, &payment); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	mtx.Lock()
	if money >= payment.USD {
		money -= payment.USD
	}

	payment.Println()

	paymentHistory = append(paymentHistory, payment)
	fmt.Println("money:", money)
	fmt.Println("payments:", paymentHistory)
	mtx.Unlock()
}

func main() {
	http.HandleFunc("/pay", payHandler)

	if err := http.ListenAndServe(":3000", nil); err != nil {
		fmt.Println(err)
	}
}
