package main

import (
	"fmt"
	"net/http"
)

func baseHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
}

func main() {
	http.HandleFunc("/tasks/", baseHandler)

	if err := http.ListenAndServe(":3000", nil); err != nil {
		fmt.Println(err)
	}

}
