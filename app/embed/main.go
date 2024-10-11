package main

import (
	"embed"
	"net/http"
)

//go:embed static/*
var static embed.FS

func main() {
	// 启动服务
	http.Handle("/static/", http.FileServer(http.FS(static)))
	http.ListenAndServe(":8082", nil)
}
