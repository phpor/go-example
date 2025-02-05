package main

import (
	"context"
	"golang.org/x/time/rate"
)

func main() {
	l := rate.NewLimiter(1, 10)
	for i := 0; i < 15; i++ {
		// 10个请求，每秒1个请求，10秒内完成
		_ = l.Wait(context.Background())
		println(i)
	}
}
