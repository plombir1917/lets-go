package main

import (
	"context"
	"fmt"
	"time"
)

type Context interface {
	Deadline() (deadline time.Time, ok bool)
	Done() <-chan struct{}
	Err() error
	Value(key any) any
}

func main() {
	WithValue()
	WithCancel()
	WithTimeout()
	WithDeadline()
}

func WithValue() {
	ctx := context.WithValue(context.Background(), "key", 123)

	fmt.Println(ctx.Value("key")) // 123
}

func WithCancel() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(2 * time.Second)
		cancel()
	}()

	<-ctx.Done()
	fmt.Println(ctx.Err()) // context cancelled
}

func WithTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	<-ctx.Done()
	fmt.Println(ctx.Err()) // context deadline exceeded
}

func WithDeadline() {
	deadline := time.Now().Add(2 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)

	defer cancel()

	<-ctx.Done()
	fmt.Println(ctx.Err())
}

// ф-ия, выполняющая какую-то медленную работу с использование контекста
// контекст - первый аргумент
func sleepRandomContext(ctx context.Context, ch chan bool) {

	// отложенное выполнение действий по очистке
	defer func() {
		fmt.Println("sleepRandomContext complete!")
		ch <- true
	}()

	// создаем канал
	sleepTimeChan := make(chan int)

	// запускаем медленное выполнение задачи в горутине
	// передаем канал для коммуникаций
	// go sleepRandom("SleepRandomContext", sleepTimeChan)

	// используем select для выхода по итсечению времени жизни контекста
	select {
	case <-ctx.Done():
		fmt.Println("Time to return")
	case sleeptime := <-sleepTimeChan:
		fmt.Println("Slept for", sleeptime, "ms")
	}
}
