package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Payment struct {
	Description string `json:"description"`
	USD         int    `json:"usd"`
	FullName    string `json:"fullName"`
	Address     string `json:"address"`
	Time        time.Time
}

type HttpResponse struct {
	Money          int       `json:"money"`
	PaymentHistory []Payment `json:"paymentHistory"`
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
	var payment Payment

	if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	payment.Time = time.Now()

	mtx.Lock()
	if money >= payment.USD {
		money -= payment.USD
	}

	payment.Println()

	paymentHistory = append(paymentHistory, payment)
	httpResponse := HttpResponse{
		Money:          money,
		PaymentHistory: paymentHistory,
	}

	b, err := json.Marshal(httpResponse)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(b); err != nil {
		fmt.Println(err)
		return
	}

	mtx.Unlock()
}

func main() {
	http.HandleFunc("/pay", payHandler)

	if err := http.ListenAndServe(":3000", nil); err != nil {
		fmt.Println(err)
	}
}
