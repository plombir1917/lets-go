package main

import (
	"context"
	"fmt"
	"time"
)

// parent
func foo(ctx context.Context, number int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Foo", number, "всё")
			return
		default:
			fmt.Println("Foo", number)
		}

		time.Sleep(100 * time.Millisecond)
	}
}

// child
func boo(ctx context.Context, number int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Boo", number, "всё")
			return
		default:
			fmt.Println("Boo", number)
		}

		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	parentContext, parentCancel := context.WithCancel(context.Background())
	childContext, childCancel := context.WithCancel(parentContext)
	go foo(parentContext, 1)
	go foo(parentContext, 2)
	go foo(parentContext, 3)

	go boo(childContext, 1)
	go boo(childContext, 2)
	go boo(childContext, 3)

	time.Sleep(1 * time.Second)

	childCancel()

	parentCancel()

	time.Sleep(3 * time.Second)

}
