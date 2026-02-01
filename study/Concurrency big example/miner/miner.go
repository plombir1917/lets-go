package miner

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func miner(
	ctx context.Context,
	wg *sync.WaitGroup,
	transferPoint chan<- int,
	n int,
	power int,
) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Я шахтер номер:", n, "мой рабочий день закончен")
			return
		case <-time.After(1 * time.Second):
			fmt.Println(" Я шахтёр номер:", n, "начал добывать уголь")
			fmt.Println("Я шахтёр номер:", n, "добыл уголь:", power)

			transferPoint <- power
			fmt.Println("Я шахтёр номер:", n, "передал уголь:", power)
		}

	}

}

func MinerPool(ctx context.Context, minerCount int) <-chan int {
	coalTransferPoint := make(chan int)

	wg := &sync.WaitGroup{}

	for i := 1; i <= minerCount; i++ {
		wg.Add(1)
		go miner(ctx, wg, coalTransferPoint, i, i*10)
	}

	go func() {
		wg.Wait()
		close(coalTransferPoint)
	}()

	return coalTransferPoint
}
