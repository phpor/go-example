package main

import (
	"context"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		select {
		case <-ctx.Done():
			println("1")
		}
	}(ctx)

	time.Sleep(3 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)
}
