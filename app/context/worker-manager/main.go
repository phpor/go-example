package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

// WorkerManager 管理多个 worker 协程
type WorkerManager struct {
	numWorkers int
	url        string
	ctx        context.Context
	cancel     context.CancelFunc
	wg         sync.WaitGroup
	eg         *errgroup.Group
}

// NewWorkerManager 创建一个新的 WorkerManager 实例
func NewWorkerManager(numWorkers int, url string) *WorkerManager {
	ctx, cancel := context.WithCancel(context.Background())
	g, _ := errgroup.WithContext(ctx)
	return &WorkerManager{
		numWorkers: numWorkers,
		url:        url,
		ctx:        ctx,
		cancel:     cancel,
		eg:         g,
	}
}

// Start 启动指定数量的 worker 协程
func (wm *WorkerManager) Start() {
	for i := 0; i < wm.numWorkers; i++ {
		wm.eg.Go(func() error {
			wm.wg.Add(1)
			defer wm.wg.Done()
			return wm.workerLoop()
		})
	}
}

// workerLoop 每个 worker 循环请求 HTTP 地址并处理返回内容
func (wm *WorkerManager) workerLoop() error {
	for {
		select {
		case <-wm.ctx.Done():
			fmt.Println("Worker shutting down gracefully")
			return nil
		default:
			resp, err := http.Get(wm.url)
			if err != nil {
				fmt.Printf("HTTP request failed: %v\n", err)
				time.Sleep(1 * time.Second) // 重试间隔
				continue
			}

			body, err := io.ReadAll(resp.Body)
			_ = resp.Body.Close()
			if err != nil {
				fmt.Printf("Failed to read response body: %v\n", err)
				time.Sleep(1 * time.Second) // 重试间隔
				continue
			}

			// 处理返回内容
			wm.processResponse(body)
		}
	}
}

// processResponse 处理 HTTP 返回的内容
func (wm *WorkerManager) processResponse(data []byte) {
	fmt.Printf("Received response: %s\n", string(data))
	// 这里可以添加更多的处理逻辑
	time.Sleep(1 * time.Second) // 模拟处理时间
}

// Shutdown 优雅地关闭所有 worker 协程
func (wm *WorkerManager) Shutdown() error {
	wm.cancel()
	if err := wm.eg.Wait(); err != nil {
		return fmt.Errorf("worker group wait: %v", err)
	}
	fmt.Println("All workers have completed their tasks")
	return nil
}

func main() {
	url := "https://example.com"
	numWorkers := 3

	manager := NewWorkerManager(numWorkers, url)
	manager.Start()

	time.Sleep(10 * time.Second) // 让 worker 运行一段时间

	if err := manager.Shutdown(); err != nil {
		fmt.Printf("Shutdown error: %v\n", err)
	}
}
