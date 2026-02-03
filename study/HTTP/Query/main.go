package main

import (
	"fmt"
	"net/http"
)

func baseHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Query().Get("v"))
}

func main() {
	http.HandleFunc("/base", baseHandler)
	if err := http.ListenAndServe(":3000", nil); err != nil {
		fmt.Println(err)
	}

}
