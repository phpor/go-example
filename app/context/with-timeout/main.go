package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	s := time.Now()
	// 创建一个带有超时功能的上下文
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	//defer cancel()

	// 启动一个goroutine执行耗时操作
	longRunningOperation(ctx)

	// 等待一段时间后取消操作
	//time.Sleep(3 * time.Second)
	fmt.Printf("Operation timed out: %s", time.Now().Sub(s).String())
}

func longRunningOperation(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Operation cancelled")
			return
		default:
			// 执行一些耗时操作
			fmt.Println("Performing some long operation")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
