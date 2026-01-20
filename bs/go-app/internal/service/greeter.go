package service

import (
	"context"
	"time"
)

func Greet(ctx context.Context, name string) (string, error) {
	select {
	case <-time.After(2 * time.Second):
		return "Hello, " + name, nil
	case <-ctx.Done():
		return "", ctx.Err()
	}
}
