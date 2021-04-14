package main

import (
	"fmt"
	"net"
	"time"
)

// golang 监听随机端口
func main() {
	listen, err := net.Listen("tcp", ":0")
	if err != nil {
		fmt.Println("listen error: ", err)
		return
	}
	time.Sleep(100 * time.Second)
	println(listen.Addr().String())
}
