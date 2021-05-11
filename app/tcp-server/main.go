package main

import (
	"fmt"
	"net"
	"time"
)

// golang 监听随机端口
func main() {
	listen, err := net.Listen("tcp", ":9090")
	if err != nil {
		fmt.Println("listen error: ", err)
		return
	}

	go func() {
		time.Sleep(3 * time.Second)
		listen.Close()
	}()
	c, err := listen.Accept()
	if err != nil {
		println(err.Error())
	} else {
		println(c.RemoteAddr())
	}
	//time.Sleep(100 * time.Second)
	println(listen.Addr().String())
}
