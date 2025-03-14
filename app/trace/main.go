package main

import (
	"context"
	"os"
	"runtime/trace"
	"time"
)

func main() {
	// 创建trace文件
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// 开始追踪
	if err := trace.Start(f); err != nil {
		panic(err)
	}
	defer trace.Stop()

	// 模拟一些工作
	for i := 0; i < 10; i++ {
		doSomeWork()
	}
}

func doSomeWork() {
	// 使用runtime/trace的Push/Pop方法来记录事件
	ctx := context.Background()
	ctx, task := trace.NewTask(ctx, "Doing some work")
	defer task.End()
	trace.WithRegion(ctx, "region1", func() {
		time.Sleep(10 * time.Millisecond)
	})
	trace.WithRegion(ctx, "region2", func() {
		time.Sleep(10 * time.Millisecond)
	})

	// 执行一些任务
	time.Sleep(time.Second)
}
