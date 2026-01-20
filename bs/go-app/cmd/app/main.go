package main

import (
	"context"
	"fmt"
	"time"

	"example.com/go-app/internal/service"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)

	defer cancel()

	msg, err := service.Greet(ctx, "Go")

	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println(msg)
}
