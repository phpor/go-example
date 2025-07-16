package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	// 创建静态文件服务
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	// SSE事件流端点
	http.HandleFunc("/events", func(w http.ResponseWriter, r *http.Request) {
		// 设置SSE响应头
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		// 每2秒发送时间戳事件
		ticker := time.NewTicker(2 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case t := <-ticker.C:
				// 发送SSE格式数据
				fmt.Fprintf(w, "data: %s\n\n", t.Format(time.RFC3339))
				w.(http.Flusher).Flush()
			case <-r.Context().Done():
				return
			}
		}
	})

	log.Println("SSE服务启动，监听: http://localhost:8080/ ...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
