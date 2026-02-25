package context

import (
	"context"
	"fmt"
	"time"
)

func longTask(ctx context.Context) {
	time.Sleep(5 * time.Second)
	fmt.Println("КОНЧАЮ")
}

func Execute() {
	ctx, cancel := context.WithCancel(context.Background())

	go longTask(ctx)

	time.Sleep(2 * time.Second)

	cancel()

	select {
	case <-ctx.Done():
		fmt.Println("Контекст завершен")
	default:
		fmt.Println("Дефолт")
	}

}
