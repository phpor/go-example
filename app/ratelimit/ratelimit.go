package main

import (
	"context"
	"golang.org/x/time/rate"
	"time"
)

func main() {
	l := rate.NewLimiter(1, 10)
	for i := 0; i < 15; i++ {
		// 10个请求，每秒1个请求，10秒内完成
		_ = l.Wait(context.Background())
		println(i)
	}
	time.Sleep(3 * time.Second)
	for i := 0; i < 3; i++ {
		_ = l.Wait(context.Background())
		println(i)
	}
}
